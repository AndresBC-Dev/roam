package repository

import (
	"github.com/AndresBC-Dev/tickets-api/v1/model"
)

type EventRepository struct {
	db     any
	Events []*model.Event
}

func (r *EventRepository) CreateOne(event model.Event) (*model.Event, error) {

	var Events = append(r.Events, &event)
	index := len(Events) - 1
	return Events[index], nil
}

func (r *EventRepository) GetOne(eventID uint) (*model.Event, error) {
	return nil, nil
}

func (r *EventRepository) GetMany(event model.Event) ([]*model.Event, error) {

	return r.Events, nil
}

func NewEventRepository(db any) *EventRepository {
	return &EventRepository{
		db: db,
	}
}
