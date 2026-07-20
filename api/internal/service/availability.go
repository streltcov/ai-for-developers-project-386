package service

import (
	"time"

	"call-booking/internal/domain"
	"call-booking/internal/repository"
)

type AvailabilityService struct {
	repo *repository.Repository
}

func NewAvailabilityService(repo *repository.Repository) *AvailabilityService {
	return &AvailabilityService{repo: repo}
}

func (s *AvailabilityService) GetAvailableSlots(eventTypeID int32, from, to time.Time) ([]domain.AvailableSlot, error) {
	et, err := s.repo.GetEventType(eventTypeID)
	if err != nil {
		return nil, err
	}

	existing, err := s.repo.GetBookingsByEventAndRange(eventTypeID, from, to)
	if err != nil {
		return nil, err
	}

	return computeSlots(from, to, et.Duration, existing), nil
}

func computeSlots(from, to time.Time, durationMin int32, existing []domain.Booking) []domain.AvailableSlot {
	duration := time.Duration(durationMin) * time.Minute
	var slots []domain.AvailableSlot

	for start := from; start.Add(duration).Before(to) || start.Add(duration).Equal(to); start = start.Add(duration) {
		end := start.Add(duration)
		if isAvailable(start, end, existing) {
			slots = append(slots, domain.AvailableSlot{
				StartTime: start,
				EndTime:   end,
			})
		}
	}

	return slots
}

func isAvailable(start, end time.Time, existing []domain.Booking) bool {
	for _, b := range existing {
		if start.Before(b.EndTime) && end.After(b.StartTime) {
			return false
		}
	}
	return true
}
