package repositories

import (
	"context"
	"time"

	"github.com/Dostonlv/booking-app/models"
)

type eventRepository struct {
	db any
}

func NewEventRepository(db any) models.EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "1",
		Name:      "Event 1",
		Location:  "Location 1",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *eventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *eventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}
