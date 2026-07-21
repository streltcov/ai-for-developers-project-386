package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"call-booking/internal/domain"
)

func (r *Repository) CreateEventType(et domain.EventType) (domain.EventType, error) {
	result, err := r.db.Exec(
		"INSERT INTO event_types (name, description, duration) VALUES (?, ?, ?)",
		et.Name, et.Description, et.Duration,
	)
	if err != nil {
		return domain.EventType{}, fmt.Errorf("insert event_type: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.EventType{}, fmt.Errorf("last insert id: %w", err)
	}

	et.ID = int32(id)
	return et, nil
}

func (r *Repository) GetEventType(id int32) (domain.EventType, error) {
	var et domain.EventType
	err := r.db.QueryRow(
		"SELECT id, name, description, duration FROM event_types WHERE id = ?", id,
	).Scan(&et.ID, &et.Name, &et.Description, &et.Duration)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.EventType{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.EventType{}, fmt.Errorf("get event_type: %w", err)
	}
	return et, nil
}

func (r *Repository) ListEventTypes() ([]domain.EventType, error) {
	rows, err := r.db.Query("SELECT id, name, description, duration FROM event_types ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("list event_types: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var eventTypes []domain.EventType
	for rows.Next() {
		var et domain.EventType
		if err := rows.Scan(&et.ID, &et.Name, &et.Description, &et.Duration); err != nil {
			return nil, fmt.Errorf("scan event_type: %w", err)
		}
		eventTypes = append(eventTypes, et)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}
	return eventTypes, nil
}

func (r *Repository) UpdateEventType(id int32, et domain.EventType) (domain.EventType, error) {
	result, err := r.db.Exec(
		"UPDATE event_types SET name = ?, description = ?, duration = ? WHERE id = ?",
		et.Name, et.Description, et.Duration, id,
	)
	if err != nil {
		return domain.EventType{}, fmt.Errorf("update event_type: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return domain.EventType{}, fmt.Errorf("rows affected: %w", err)
	}
	if rows == 0 {
		return domain.EventType{}, domain.ErrNotFound
	}

	et.ID = id
	return et, nil
}

func (r *Repository) DeleteEventType(id int32) error {
	var count int
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM bookings WHERE event_type_id = ?", id,
	).Scan(&count)
	if err != nil {
		return fmt.Errorf("count bookings: %w", err)
	}
	if count > 0 {
		return domain.ErrSlotTaken
	}

	result, err := r.db.Exec("DELETE FROM event_types WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete event_type: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rows == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *Repository) EventTypeExists(id int32) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM event_types WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check event_type existence: %w", err)
	}
	return exists, nil
}
