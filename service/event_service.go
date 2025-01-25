package service

import (
	"errors"
	"project-akhir-exam/entity"
	"project-akhir-exam/repository"

	"github.com/gosimple/slug"
)

type GetEventDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateEventInput struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Price       int    `form:"price" binding:"required"`
	Stock       int    `form:"stock" binding:"required"`
}

type EventService interface {
	GetAllEvents() ([]entity.Event, error)
	GetEventByID(input GetEventDetailInput) (entity.Event, error)
	CreateEvent(input CreateEventInput, fileLocation string) (entity.Event, error)
	UpdateEvent(inputID GetEventDetailInput, inputData CreateEventInput) (entity.Event, error)
	DeleteEvent(inputID GetEventDetailInput) (entity.Event, error)
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) *eventService {
	return &eventService{eventRepository}
}

func (s *eventService) GetAllEvents() ([]entity.Event, error) {
	events, err := s.eventRepository.FindAll()
	if err != nil {
		return events, err
	}
	return events, nil
}

func (s *eventService) GetEventByID(input GetEventDetailInput) (entity.Event, error) {
	event, err := s.eventRepository.FindByID(input.ID)

	if err != nil {
		return event, err
	}

	if event.ID == 0 {
		return event, errors.New("no event found on that id")
	}

	return event, nil
}

func (s *eventService) CreateEvent(input CreateEventInput, fileLocation string) (entity.Event, error) {
	event := entity.Event{}
	event.Name = input.Name
	event.Description = input.Description
	event.Price = input.Price
	event.Stock = input.Stock
	event.Slug = slug.Make(input.Name)
	event.Image = fileLocation

	newEvent, err := s.eventRepository.Save(event)
	if err != nil {
		return newEvent, err
	}

	return newEvent, nil
}

func (s *eventService) UpdateEvent(inputID GetEventDetailInput, inputData CreateEventInput) (entity.Event, error) {
	event, err := s.eventRepository.FindByID(inputID.ID)
	if err != nil {
		return event, err
	}

	if event.ID == 0 {
		return event, errors.New("no event found on that id")
	}

	event.Name = inputData.Name
	event.Description = inputData.Description
	event.Price = inputData.Price
	event.Stock = inputData.Stock
	event.Slug = slug.Make(inputData.Name)

	updatedEvent, err := s.eventRepository.Update(event)
	if err != nil {
		return updatedEvent, err
	}

	return updatedEvent, nil
}

func (s *eventService) DeleteEvent(inputID GetEventDetailInput) (entity.Event, error) {
	event, err := s.eventRepository.FindByID(inputID.ID)
	if err != nil {
		return event, err
	}

	if event.ID == 0 {
		return event, errors.New("no event found on that id")
	}

	deleteEvent, err := s.eventRepository.Delete(event)
	if err != nil {
		return deleteEvent, err
	}

	return deleteEvent, nil
}
