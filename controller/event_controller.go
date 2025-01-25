package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"project-akhir-exam/helper"
	"project-akhir-exam/service"

	"github.com/gin-gonic/gin"
)

type eventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) *eventController {
	return &eventController{eventService}
}

func (h *eventController) GetAllEvents(c *gin.Context) {
	events, err := h.eventService.GetAllEvents()
	if err != nil {
		response := helper.APIResponse("Error to get events", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of events", http.StatusOK, "success", helper.FormatEvents(events))
	c.JSON(http.StatusOK, response)
}

func (h *eventController) GetEvent(c *gin.Context) {
	var input service.GetEventDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of event", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	eventDetail, err := h.eventService.GetEventByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail of event", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Event detail", http.StatusOK, "success", helper.FormatEventDetail(eventDetail))
	c.JSON(http.StatusOK, response)
}

func (h *eventController) CreateEvent(c *gin.Context) {
	var input service.CreateEventInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create event", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("public/event/%d-%s", rand.Int(), file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newEvent, err := h.eventService.CreateEvent(input, path)
	if err != nil {
		response := helper.APIResponse("Failed to create event", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create event", http.StatusOK, "success", helper.FormatEvent(newEvent))
	c.JSON(http.StatusOK, response)
}

func (h *eventController) UpdateEvent(c *gin.Context) {
	var inputID service.GetEventDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update event", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData service.CreateEventInput

	err = c.ShouldBind(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update event", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedEvent, err := h.eventService.UpdateEvent(inputID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to update event", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update event", http.StatusOK, "success", helper.FormatEvent(updatedEvent))
	c.JSON(http.StatusOK, response)
}

func (h *eventController) DeleteEvent(c *gin.Context) {
	var input service.GetEventDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to delete event", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deletedEvent, err := h.eventService.DeleteEvent(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to delete event", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete event", http.StatusOK, "success", helper.FormatEventDetail(deletedEvent))
	c.JSON(http.StatusOK, response)
}
