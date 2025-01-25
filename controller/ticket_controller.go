package controller

import (
	"net/http"
	"project-akhir-exam/entity"
	"project-akhir-exam/helper"
	"project-akhir-exam/service"

	"github.com/gin-gonic/gin"
)

type ticketController struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) *ticketController {
	return &ticketController{ticketService}
}

func (h *ticketController) GetEventTickets(c *gin.Context) {
	var input service.GetEventTicketInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get event's tickets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tickets, err := h.ticketService.GetTicketByEventID(input)
	if err != nil {
		// errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get event's tickets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Event's tickets", http.StatusOK, "success", helper.FormatEventTickets(tickets))
	c.JSON(http.StatusOK, response)
}

func (h *ticketController) GetUserTickets(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(entity.User)
	userID := currentUser.ID

	tickets, err := h.ticketService.GetTicketByUserID(userID)
	if err != nil {
		// errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get user's tickets", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's tickets", http.StatusOK, "success", helper.FormatUserTickets(tickets))
	c.JSON(http.StatusOK, response)
}

func (h *ticketController) GetTicket(c *gin.Context) {
	var input service.GetTicketInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of ticket", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	ticketDetail, err := h.ticketService.GetTicketByID(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to get detail of ticket", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Ticket detail", http.StatusOK, "success", helper.FormatCreateTicket(ticketDetail))
	c.JSON(http.StatusOK, response)
}

func (h *ticketController) CreateTicket(c *gin.Context) {
	var input service.CreateTicketInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create ticket", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(entity.User)

	input.User = currentUser
	// input.Event = inputID

	newTicket, err := h.ticketService.CreateTicket(input)
	if err != nil {
		response := helper.APIResponse("Failed to create ticket", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create ticket", http.StatusOK, "success", helper.FormatCreateTicket(newTicket))
	c.JSON(http.StatusOK, response)
}

func (h *ticketController) MarkPaid(c *gin.Context) {
	var inputID service.GetTicketInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to pay ticket", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	markPaid, err := h.ticketService.MarkPaid(inputID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to pay ticket", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to pay ticket", http.StatusOK, "success", helper.FormatCreateTicket(markPaid))
	c.JSON(http.StatusOK, response)
}

func (h *ticketController) MarkCancel(c *gin.Context) {
	var inputID service.GetTicketInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to cancel ticket", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	markCancel, err := h.ticketService.MarkCancel(inputID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to cancel ticket", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to cancel ticket", http.StatusOK, "success", helper.FormatCreateTicket(markCancel))
	c.JSON(http.StatusOK, response)
}

// func (h *ticketController) MarkStatus(c *gin.Context) {
// 	var inputID service.GetTicketInput

// 	err := c.ShouldBindUri(&inputID)
// 	if err != nil {
// 		response := helper.APIResponse("Failed to change status", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	var inputData service.GetTicketStatusInput
// 	err = c.ShouldBindJSON(&inputData)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}

// 		response := helper.APIResponse("Failed to change status", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	markStatus, err := h.ticketService.MarkStatus(inputID, inputData)
// 	if err != nil {
// 		errorMessage := gin.H{"errors": err.Error()}
// 		response := helper.APIResponse("Failed to change status", http.StatusBadRequest, "error", errorMessage)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse("Success to change status", http.StatusOK, "success", helper.FormatCreateTicket(markStatus))
// 	c.JSON(http.StatusOK, response)
// }
