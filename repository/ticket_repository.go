package repository

import (
	"project-akhir-exam/entity"

	"gorm.io/gorm"
)

type TicketRepository interface {
	GetByEventID(EventID int) ([]entity.Ticket, error)
	GetByUserID(UserID int) ([]entity.Ticket, error)
	FindByID(ID int) (entity.Ticket, error)
	Save(ticket entity.Ticket) (entity.Ticket, error)
	MarkStatus(ticket entity.Ticket) (entity.Ticket, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *ticketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) GetByEventID(EventID int) ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	err := r.db.Preload("User").Where("event_id = ?", EventID).Order("id desc").Find(&tickets).Error
	// err := r.db.Preload("User").Preload("Product").Where("product_id = ?", productID).Order("id desc").Find(&tickets).Error
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

func (r *ticketRepository) GetByUserID(UserID int) ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	err := r.db.Preload("Event").Preload("User").Where("user_id = ?", UserID).Order("id desc").Find(&tickets).Error
	if err != nil {
		return tickets, err
	}

	return tickets, nil
}

func (r *ticketRepository) FindByID(ID int) (entity.Ticket, error) {
	var ticket entity.Ticket

	err := r.db.Where("id = ?", ID).Find(&ticket).Error
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (r *ticketRepository) Save(ticket entity.Ticket) (entity.Ticket, error) {
	var event entity.Event
	r.db.First(&event, ticket.EventID)

	ticket.Total = event.Price * ticket.Quantity
	err := r.db.Preload("Event").Preload("User").Create(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (r *ticketRepository) MarkStatus(ticket entity.Ticket) (entity.Ticket, error) {
	err := r.db.Save(&ticket).Error
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}
