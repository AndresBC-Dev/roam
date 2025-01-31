package model

import (
	"time"
)

type Event struct {
	ID        uint
	Name      string
	Location  string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EventRepository interface {
	CreateOne(event Event) (*Event, error)
	GetOne(eventID uint) (*Event, error)
	GetMany(event Event) ([]*Event, error)
}
