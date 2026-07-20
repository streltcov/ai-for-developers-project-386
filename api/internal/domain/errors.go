package domain

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrSlotTaken  = errors.New("slot already booked")
	ErrValidation = errors.New("validation error")
)
