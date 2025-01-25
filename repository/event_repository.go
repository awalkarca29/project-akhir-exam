package repository

import (
	"project-akhir-exam/entity"

	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	FindByID(ID int) (entity.Event, error)
	Save(event entity.Event) (entity.Event, error)
	Update(event entity.Event) (entity.Event, error)
	Delete(event entity.Event) (entity.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) FindAll() ([]entity.Event, error) {
	var events []entity.Event

	err := r.db.Find(&events).Error
	if err != nil {
		return events, err
	}

	return events, nil
}

func (r *eventRepository) FindByID(ID int) (entity.Event, error) {
	var event entity.Event

	err := r.db.Where("id = ?", ID).Find(&event).Error
	if err != nil {
		return event, err
	}
	return event, nil
}

func (r *eventRepository) Save(event entity.Event) (entity.Event, error) {
	err := r.db.Create(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *eventRepository) Update(event entity.Event) (entity.Event, error) {
	err := r.db.Save(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *eventRepository) Delete(event entity.Event) (entity.Event, error) {
	err := r.db.Delete(&event).Error
	if err != nil {
		return event, err
	}

	return event, nil
}
