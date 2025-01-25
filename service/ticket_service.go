package service

import (
	"errors"
	"project-akhir-exam/entity"
	"project-akhir-exam/repository"
)

type GetEventTicketInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetTicketInput struct {
	ID int `uri:"id" binding:"required"`
}

type GetTicketStatusInput struct {
	Status string `json:"status" binding:"required"`
}

type CreateTicketInput struct {
	EventID       int    `json:"event_id" binding:"required"`
	Quantity      int    `json:"quantity" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	User          entity.User
}

type TicketService interface {
	GetTicketByEventID(input GetEventTicketInput) ([]entity.Ticket, error)
	GetTicketByUserID(userID int) ([]entity.Ticket, error)
	GetTicketByID(input GetTicketInput) (entity.Ticket, error)
	CreateTicket(input CreateTicketInput) (entity.Ticket, error)
	MarkPaid(input GetTicketInput) (entity.Ticket, error)
	MarkCancel(input GetTicketInput) (entity.Ticket, error)
	// MarkStatus(input GetTicketInput, inputData GetTicketStatusInput) (entity.Ticket, error)
}

type ticketService struct {
	ticketRepository repository.TicketRepository
}

func NewTicketService(ticketRepository repository.TicketRepository) *ticketService {
	return &ticketService{ticketRepository}
}

func (s *ticketService) GetTicketByEventID(input GetEventTicketInput) ([]entity.Ticket, error) {
	tickets, err := s.ticketRepository.GetByEventID(input.ID)
	if err != nil {
		return tickets, err
	}

	// if tickets[0].EventID == 0 {
	// 	return tickets, errors.New("no ticket found on that event id")
	// }

	return tickets, nil
}

func (s *ticketService) GetTicketByUserID(userID int) ([]entity.Ticket, error) {
	tickets, err := s.ticketRepository.GetByUserID(userID)
	if err != nil {
		return tickets, err
	}

	// if tickets[0].UserID == 0 {
	// 	return tickets, errors.New("no ticket found on that user id")
	// }

	return tickets, nil
}

func (s *ticketService) GetTicketByID(input GetTicketInput) (entity.Ticket, error) {
	ticket, err := s.ticketRepository.FindByID(input.ID)

	if err != nil {
		return ticket, err
	}

	if ticket.ID == 0 {
		return ticket, errors.New("no ticket found on that id")
	}

	return ticket, nil
}

func (s *ticketService) CreateTicket(input CreateTicketInput) (entity.Ticket, error) {
	ticket := entity.Ticket{}
	ticket.UserID = input.User.ID
	ticket.EventID = input.EventID
	ticket.Quantity = input.Quantity
	ticket.PaymentMethod = input.PaymentMethod
	ticket.Status = "pending"

	newTicket, err := s.ticketRepository.Save(ticket)
	if err != nil {
		return newTicket, err
	}

	return newTicket, nil
}

func (s *ticketService) MarkPaid(input GetTicketInput) (entity.Ticket, error) {
	ticket, err := s.ticketRepository.FindByID(input.ID)
	if err != nil {
		return ticket, err
	}

	if ticket.ID == 0 {
		return ticket, errors.New("no ticket found on that id")
	}

	ticket.Status = "paid"

	markPaid, err := s.ticketRepository.MarkStatus(ticket)
	if err != nil {
		return markPaid, err
	}

	return markPaid, nil
}

func (s *ticketService) MarkCancel(input GetTicketInput) (entity.Ticket, error) {
	ticket, err := s.ticketRepository.FindByID(input.ID)
	if err != nil {
		return ticket, err
	}

	if ticket.ID == 0 {
		return ticket, errors.New("no ticket found on that id")
	}

	ticket.Status = "cancel"

	markCancel, err := s.ticketRepository.MarkStatus(ticket)
	if err != nil {
		return markCancel, err
	}

	return markCancel, nil
}

// func (s *ticketService) MarkStatus(input GetTicketInput, inputData GetTicketStatusInput) (entity.Ticket, error) {
// 	ticket, err := s.ticketRepository.FindByID(input.ID)
// 	if err != nil {
// 		return ticket, err
// 	}

// 	if ticket.ID == 0 {
// 		return ticket, errors.New("no ticket found on that id")
// 	}

// 	ticket.Status = inputData.Status

// 	markStatus, err := s.ticketRepository.MarkStatus(ticket)
// 	if err != nil {
// 		return markStatus, err
// 	}

// 	return markStatus, nil
// }
