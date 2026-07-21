package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"call-booking/internal/domain"
)

func (h *Handler) publicListEventTypes(w http.ResponseWriter, r *http.Request) {
	eventTypes, err := h.repo.ListEventTypes()
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}
	writeJSON(w, http.StatusOK, eventTypes)
}

func (h *Handler) getAvailability(w http.ResponseWriter, r *http.Request) {
	eventTypeID := r.URL.Query().Get("eventTypeId")
	if eventTypeID == "" {
		writeError(w, http.StatusBadRequest, "eventTypeId is required", "VALIDATION_ERROR")
		return
	}

	var id int32
	if _, err := parseQueryParam(r, "eventTypeId", &id); err != nil {
		writeError(w, http.StatusBadRequest, "invalid eventTypeId", "VALIDATION_ERROR")
		return
	}

	now := time.Now().UTC()
	from := now
	to := now.AddDate(0, 0, 14)

	if fromStr := r.URL.Query().Get("from"); fromStr != "" {
		parsed, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid from date, use YYYY-MM-DD", "VALIDATION_ERROR")
			return
		}
		from = parsed
	}

	if toStr := r.URL.Query().Get("to"); toStr != "" {
		parsed, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid to date, use YYYY-MM-DD", "VALIDATION_ERROR")
			return
		}
		to = parsed
	}

	slots, err := h.availability.GetAvailableSlots(id, from, to)
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	et, err := h.repo.GetEventType(id)
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusOK, domain.AvailabilityResponse{
		EventTypeID: et.ID,
		Duration:    et.Duration,
		Slots:       slots,
	})
}

func (h *Handler) createBooking(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EventTypeID int32   `json:"eventTypeId"`
		StartTime   string  `json:"startTime"`
		GuestName   string  `json:"guestName"`
		GuestEmail  string  `json:"guestEmail"`
		Notes       *string `json:"notes,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body", "VALIDATION_ERROR")
		return
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid startTime, use RFC3339 format", "VALIDATION_ERROR")
		return
	}

	booking, err := h.booking.CreateBooking(
		req.EventTypeID,
		startTime,
		req.GuestName,
		req.GuestEmail,
		req.Notes,
	)
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusCreated, booking)
}

func parseQueryParam(r *http.Request, name string, dest any) (string, error) {
	val := r.URL.Query().Get(name)
	if val == "" {
		return "", nil
	}
	switch d := dest.(type) {
	case *int32:
		var n int
		for _, c := range val {
			if c < '0' || c > '9' {
				return val, domain.ErrValidation
			}
			n = n*10 + int(c-'0')
		}
		*d = int32(n)
	}
	return val, nil
}
