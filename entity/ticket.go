package entity

import (
	"time"
)

type Ticket struct {
	ID            int
	UserID        int
	EventID       int
	Quantity      int
	Total         int
	PaymentMethod string
	Status        string
	User          User  `gorm:"foreignKey:UserID"`
	Event         Event `gorm:"foreignKey:EventID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
