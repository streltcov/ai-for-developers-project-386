package domain

import "time"

type EventType struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Duration    int32   `json:"duration"`
}

type Booking struct {
	ID          int32     `json:"id"`
	EventTypeID int32     `json:"eventTypeId"`
	EventType   EventType `json:"eventType"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	GuestName   string    `json:"guestName"`
	GuestEmail  string    `json:"guestEmail"`
	Notes       *string   `json:"notes,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

type AvailableSlot struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
