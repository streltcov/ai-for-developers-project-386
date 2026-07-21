package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"call-booking/internal/domain"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) ownerListEventTypes(w http.ResponseWriter, r *http.Request) {
	eventTypes, err := h.repo.ListEventTypes()
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}
	writeJSON(w, http.StatusOK, eventTypes)
}

func (h *Handler) createEventType(w http.ResponseWriter, r *http.Request) {
	var et domain.EventType
	if err := json.NewDecoder(r.Body).Decode(&et); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body", "VALIDATION_ERROR")
		return
	}

	if et.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required", "VALIDATION_ERROR")
		return
	}

	if et.Duration <= 0 {
		writeError(w, http.StatusBadRequest, "duration must be positive", "VALIDATION_ERROR")
		return
	}

	created, err := h.repo.CreateEventType(et)
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

func (h *Handler) getEventType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", "VALIDATION_ERROR")
		return
	}

	et, err := h.repo.GetEventType(int32(id))
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusOK, et)
}

func (h *Handler) updateEventType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", "VALIDATION_ERROR")
		return
	}

	var et domain.EventType
	if err := json.NewDecoder(r.Body).Decode(&et); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body", "VALIDATION_ERROR")
		return
	}

	updated, err := h.repo.UpdateEventType(int32(id), et)
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

func (h *Handler) deleteEventType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id", "VALIDATION_ERROR")
		return
	}

	if err := h.repo.DeleteEventType(int32(id)); err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusNoContent, nil)
}

func (h *Handler) listBookings(w http.ResponseWriter, r *http.Request) {
	var from, to *time.Time
	var eventTypeID *int32

	if fromStr := r.URL.Query().Get("from"); fromStr != "" {
		parsed, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid from date, use YYYY-MM-DD", "VALIDATION_ERROR")
			return
		}
		from = &parsed
	}

	if toStr := r.URL.Query().Get("to"); toStr != "" {
		parsed, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid to date, use YYYY-MM-DD", "VALIDATION_ERROR")
			return
		}
		to = &parsed
	}

	if idStr := r.URL.Query().Get("eventTypeId"); idStr != "" {
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid eventTypeId", "VALIDATION_ERROR")
			return
		}
		eventTypeID = &[]int32{int32(id)}[0]
	}

	bookings, err := h.repo.ListBookings(from, to, eventTypeID)
	if err != nil {
		status, message, code := mapError(err)
		writeError(w, status, message, code)
		return
	}

	writeJSON(w, http.StatusOK, bookings)
}
