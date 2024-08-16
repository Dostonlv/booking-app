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
	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *eventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(&models.Event{}).Where("id = ?", eventId).First(&event)
	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *eventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Create(event)

	if res.Error != nil {
		return nil, res.Error
	}

	return event, nil
}

func (r *eventRepository) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(&models.Event{}).Where("id = ?", eventId).Updates(updateData)
	if res.Error != nil {
		return nil, res.Error
	}

	getRes := r.db.Where("id = ?", eventId).First(&event)
	if res.Error != nil {
		return nil, getRes.Error
	}

	return event, nil
}

func (r *eventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	res := r.db.Delete(&models.Event{}, eventId)
	if res.Error != nil {
		return res.Error
	}

	return res.Error
}
