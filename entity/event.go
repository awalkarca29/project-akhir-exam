package entity

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Slug        string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tickets     []Ticket
}
