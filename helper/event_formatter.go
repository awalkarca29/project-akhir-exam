package helper

import "project-akhir-exam/entity"

type EventFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

func FormatEvent(event entity.Event) EventFormatter {
	eventFormatter := EventFormatter{}
	eventFormatter.ID = event.ID
	eventFormatter.Name = event.Name
	eventFormatter.Description = event.Description
	eventFormatter.Slug = event.Slug
	eventFormatter.Price = event.Price
	eventFormatter.Stock = event.Stock
	eventFormatter.Image = event.Image

	return eventFormatter
}

func FormatEvents(events []entity.Event) []EventFormatter {
	eventsFormatter := []EventFormatter{}

	for _, event := range events {
		eventFormatter := FormatEvent(event)
		eventsFormatter = append(eventsFormatter, eventFormatter)
	}

	return eventsFormatter
}

type EventDetailFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Price       int    `json:"price"`
	Stock       int    `json:"int"`
	Image       string `json:"image"`
}

func FormatEventDetail(event entity.Event) EventDetailFormatter {
	eventDetailFormatter := EventDetailFormatter{}
	eventDetailFormatter.ID = event.ID
	eventDetailFormatter.Name = event.Name
	eventDetailFormatter.Description = event.Description
	eventDetailFormatter.Slug = event.Slug
	eventDetailFormatter.Price = event.Price
	eventDetailFormatter.Stock = event.Stock
	eventDetailFormatter.Image = event.Image

	return eventDetailFormatter
}
