package repositories

import (
	"context"

	"github.com/Dostonlv/booking-app/models"
	"gorm.io/gorm"
)

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (r *ticketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}
	res := r.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets)
	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *ticketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Where("id = ?", ticketId).First(&ticket)
	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *ticketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, ticket.ID)
}

func (r *ticketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)
	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, ticketId)
}
