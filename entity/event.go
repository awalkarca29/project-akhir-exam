package entity

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Slug        string
	Image       string
	Price       int
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tickets     []Ticket
}
