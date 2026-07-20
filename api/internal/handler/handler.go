package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"call-booking/internal/domain"
	"call-booking/internal/repository"
	"call-booking/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	repo         *repository.Repository
	availability *service.AvailabilityService
	booking      *service.BookingService
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo:         repo,
		availability: service.NewAvailabilityService(repo),
		booking:      service.NewBookingService(repo),
	}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Route("/api/public", func(r chi.Router) {
		r.Get("/event-types", h.publicListEventTypes)
		r.Get("/availability", h.getAvailability)
		r.Post("/bookings", h.createBooking)
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/event-types", h.ownerListEventTypes)
		r.Post("/event-types", h.createEventType)
		r.Get("/event-types/{id}", h.getEventType)
		r.Put("/event-types/{id}", h.updateEventType)
		r.Delete("/event-types/{id}", h.deleteEventType)
		r.Get("/bookings", h.listBookings)
	})

	return r
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func writeError(w http.ResponseWriter, status int, message string, code string) {
	writeJSON(w, status, map[string]string{
		"error": message,
		"code":  code,
	})
}

func mapError(err error) (int, string, string) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound, err.Error(), "NOT_FOUND"
	case errors.Is(err, domain.ErrSlotTaken):
		return http.StatusConflict, err.Error(), "SLOT_TAKEN"
	case errors.Is(err, domain.ErrValidation):
		return http.StatusBadRequest, err.Error(), "VALIDATION_ERROR"
	default:
		return http.StatusInternalServerError, "internal server error", "INTERNAL_ERROR"
	}
}
