package repositories

import (
	"context"

	"github.com/Dostonlv/booking-app/models"
	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}
	res := r.db.Model(&models.Event{}).Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *eventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *eventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}
