package service

import (
	"fmt"
	"strings"
	"time"

	"call-booking/internal/domain"
	"call-booking/internal/repository"
)

type BookingService struct {
	repo *repository.Repository
}

func NewBookingService(repo *repository.Repository) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) CreateBooking(
	eventTypeID int32,
	startTime time.Time,
	guestName, guestEmail string,
	notes *string,
) (domain.Booking, error) {
	if err := validateBookingInput(eventTypeID, startTime, guestName, guestEmail); err != nil {
		return domain.Booking{}, err
	}

	et, err := s.repo.GetEventType(eventTypeID)
	if err != nil {
		return domain.Booking{}, err
	}

	endTime := startTime.Add(time.Duration(et.Duration) * time.Minute)

	overlap, err := s.repo.HasOverlappingBooking(eventTypeID, startTime, endTime)
	if err != nil {
		return domain.Booking{}, fmt.Errorf("check overlap: %w", err)
	}
	if overlap {
		return domain.Booking{}, domain.ErrSlotTaken
	}

	booking := domain.Booking{
		EventTypeID: eventTypeID,
		StartTime:   startTime,
		EndTime:     endTime,
		GuestName:   strings.TrimSpace(guestName),
		GuestEmail:  strings.TrimSpace(guestEmail),
		Notes:       notes,
	}

	return s.repo.CreateBooking(booking)
}

func validateBookingInput(eventTypeID int32, startTime time.Time, guestName, guestEmail string) error {
	if eventTypeID <= 0 {
		return fmt.Errorf("%w: event_type_id must be positive", domain.ErrValidation)
	}

	if startTime.Before(time.Now().UTC()) {
		return fmt.Errorf("%w: start_time must be in the future", domain.ErrValidation)
	}

	if strings.TrimSpace(guestName) == "" {
		return fmt.Errorf("%w: guest_name is required", domain.ErrValidation)
	}

	if strings.TrimSpace(guestEmail) == "" {
		return fmt.Errorf("%w: guest_email is required", domain.ErrValidation)
	}

	if !strings.Contains(guestEmail, "@") || !strings.Contains(guestEmail, ".") {
		return fmt.Errorf("%w: invalid email format", domain.ErrValidation)
	}

	return nil
}
