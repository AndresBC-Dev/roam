package repository

import (
	"time"

	"github.com/AndresBC-Dev/tickets-api/v1/model"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) CreateOne(event model.Event) (*model.Event, error) {
	return nil, nil
}

func (r *EventRepository) GetOne(eventID uint) (*model.Event, error) {
	return nil, nil
}

func (r *EventRepository) GetMany(event model.Event) ([]*model.Event, error) {
	events := []*model.Event{}

	events = append(events, &model.Event{
		ID:        12,
		Name:      "morat concert",
		Location:  "madrid",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}

func NewEventRepository(db any) model.EventRepository {
	return &EventRepository{
		db: db,
	}
}
