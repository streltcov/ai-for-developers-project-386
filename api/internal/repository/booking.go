package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"call-booking/internal/domain"
)

func (r *Repository) CreateBooking(b domain.Booking) (domain.Booking, error) {
	result, err := r.db.Exec(
		`INSERT INTO bookings (event_type_id, start_time, end_time, guest_name, guest_email, notes)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		b.EventTypeID, b.StartTime, b.EndTime, b.GuestName, b.GuestEmail, b.Notes,
	)
	if err != nil {
		return domain.Booking{}, fmt.Errorf("insert booking: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Booking{}, fmt.Errorf("last insert id: %w", err)
	}

	b.ID = int32(id)
	b.CreatedAt = time.Now().UTC()
	return b, nil
}

func (r *Repository) ListBookings(from, to *time.Time, eventTypeID *int32) ([]domain.Booking, error) {
	query := `
		SELECT b.id, b.event_type_id, b.start_time, b.end_time,
		       b.guest_name, b.guest_email, b.notes, b.created_at,
		       et.id, et.name, et.description, et.duration
		FROM bookings b
		JOIN event_types et ON b.event_type_id = et.id
		WHERE b.start_time >= datetime('now')
	`

	args := []any{}

	if from != nil {
		query += " AND b.start_time >= ?"
		args = append(args, from.Format(time.RFC3339))
	}

	if to != nil {
		query += " AND b.start_time <= ?"
		args = append(args, to.Format(time.RFC3339))
	}

	if eventTypeID != nil {
		query += " AND b.event_type_id = ?"
		args = append(args, *eventTypeID)
	}

	query += " ORDER BY b.start_time"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("list bookings: %w", err)
	}
	defer func() { _ = rows.Close() }()

	bookings := make([]domain.Booking, 0)
	for rows.Next() {
		var b domain.Booking
		if err := rows.Scan(
			&b.ID, &b.EventTypeID, &b.StartTime, &b.EndTime,
			&b.GuestName, &b.GuestEmail, &b.Notes, &b.CreatedAt,
			&b.EventType.ID, &b.EventType.Name, &b.EventType.Description, &b.EventType.Duration,
		); err != nil {
			return nil, fmt.Errorf("scan booking: %w", err)
		}
		bookings = append(bookings, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}
	return bookings, nil
}

func (r *Repository) GetBookingsByEventAndRange(eventTypeID int32, from, to time.Time) ([]domain.Booking, error) {
	rows, err := r.db.Query(
		`SELECT id, event_type_id, start_time, end_time, guest_name, guest_email, notes, created_at
		 FROM bookings
		 WHERE event_type_id = ? AND start_time >= ? AND end_time <= ?
		 ORDER BY start_time`,
		eventTypeID, from.Format(time.RFC3339), to.Format(time.RFC3339),
	)
	if err != nil {
		return nil, fmt.Errorf("get bookings by event and range: %w", err)
	}
	defer func() { _ = rows.Close() }()

	bookings := make([]domain.Booking, 0)
	for rows.Next() {
		var b domain.Booking
		if err := rows.Scan(
			&b.ID, &b.EventTypeID, &b.StartTime, &b.EndTime,
			&b.GuestName, &b.GuestEmail, &b.Notes, &b.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan booking: %w", err)
		}
		bookings = append(bookings, b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}
	return bookings, nil
}

func (r *Repository) HasOverlappingBooking(eventTypeID int32, startTime, endTime time.Time) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		`SELECT EXISTS(
			SELECT 1 FROM bookings
			WHERE event_type_id = ?
			  AND start_time < ? AND end_time > ?
		)`,
		eventTypeID, endTime.Format(time.RFC3339), startTime.Format(time.RFC3339),
	).Scan(&exists)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("check overlap: %w", err)
	}
	return exists, nil
}
