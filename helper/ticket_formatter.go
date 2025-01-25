package helper

import (
	"project-akhir-exam/entity"
	"time"
)

type EventTicketFormatter struct {
	ID            int       `json:"id"`
	EventID       int       `json:"event_id"`
	UserID        int       `json:"user_id"`
	Name          string    `json:"name"`
	Quantity      int       `json:"quantity"`
	Total         int       `json:"total"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

func FormatEventTicket(ticket entity.Ticket) EventTicketFormatter {
	formatter := EventTicketFormatter{}
	formatter.ID = ticket.ID
	formatter.EventID = ticket.EventID
	formatter.UserID = ticket.UserID
	formatter.Name = ticket.User.Name
	formatter.Quantity = ticket.Quantity
	formatter.Total = ticket.Total
	formatter.PaymentMethod = ticket.PaymentMethod
	formatter.Status = ticket.Status
	formatter.CreatedAt = ticket.CreatedAt
	return formatter
}

func FormatEventTickets(tickets []entity.Ticket) []EventTicketFormatter {
	if len(tickets) == 0 {
		return []EventTicketFormatter{}
	}
	var ticketsFormatter []EventTicketFormatter

	for _, ticket := range tickets {
		formatter := FormatEventTicket(ticket)
		ticketsFormatter = append(ticketsFormatter, formatter)
	}

	return ticketsFormatter
}

type UserTicketFormatter struct {
	ID            int                      `json:"id"`
	EventID       int                      `json:"event_id"`
	UserID        int                      `json:"user_id"`
	Name          string                   `json:"name"`
	Quantity      int                      `json:"quantity"`
	Total         int                      `json:"total"`
	PaymentMethod string                   `json:"payment_method"`
	Status        string                   `json:"status"`
	CreatedAt     time.Time                `json:"created_at"`
	Event         EventUserTicketFormatter `json:"event"`
}

type EventUserTicketFormatter struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func FormatUserTicket(ticket entity.Ticket) UserTicketFormatter {
	formatter := UserTicketFormatter{}
	formatter.ID = ticket.ID
	formatter.EventID = ticket.EventID
	formatter.UserID = ticket.UserID
	formatter.Name = ticket.User.Name
	formatter.Quantity = ticket.Quantity
	formatter.Total = ticket.Total
	formatter.PaymentMethod = ticket.PaymentMethod
	formatter.Status = ticket.Status
	formatter.CreatedAt = ticket.CreatedAt

	eventFormatter := EventUserTicketFormatter{}
	eventFormatter.Name = ticket.Event.Name
	eventFormatter.Image = ticket.Event.Image

	formatter.Event = eventFormatter

	return formatter
}

func FormatUserTickets(tickets []entity.Ticket) []UserTicketFormatter {
	if len(tickets) == 0 {
		return []UserTicketFormatter{}
	}
	var ticketsFormatter []UserTicketFormatter

	for _, ticket := range tickets {
		formatter := FormatUserTicket(ticket)
		ticketsFormatter = append(ticketsFormatter, formatter)
	}

	return ticketsFormatter
}

type CreateTicketFormatter struct {
	ID            int       `json:"id"`
	EventID       int       `json:"event_id"`
	UserID        int       `json:"user_id"`
	Quantity      int       `json:"quantity"`
	Total         int       `json:"total"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	// Name          string                   `json:"name"`
	// Event         EventUserTicketFormatter `json:"event"`
}

func FormatCreateTicket(ticket entity.Ticket) CreateTicketFormatter {
	formatter := CreateTicketFormatter{}
	formatter.ID = ticket.ID
	formatter.EventID = ticket.EventID
	formatter.UserID = ticket.UserID
	// formatter.Name = ticket.User.Name
	formatter.Quantity = ticket.Quantity
	formatter.Total = ticket.Total
	formatter.PaymentMethod = ticket.PaymentMethod
	formatter.Status = ticket.Status
	formatter.CreatedAt = ticket.CreatedAt

	// eventFormatter := EventUserTicketFormatter{}
	// eventFormatter.Name = ticket.Event.Name
	// eventFormatter.Image = ticket.Event.Image

	// formatter.Event = eventFormatter

	return formatter
}
