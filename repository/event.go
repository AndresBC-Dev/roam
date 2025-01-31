package repository

import (
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
	return nil, nil
}

func NewEventRepository(db any) model.EventRepository {
	return &EventRepository{
		db: db,
	}
}
