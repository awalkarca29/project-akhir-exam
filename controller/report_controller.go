package controller

import (
	"log"
	"net/http"
	"project-akhir-exam/helper"
	"project-akhir-exam/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController interface {
	SummaryHandler(c *gin.Context)
	EventHandler(c *gin.Context)
}

type reportController struct {
	reportService service.ReportService
}

func NewReportController(reportService service.ReportService) ReportController {
	return &reportController{reportService: reportService}
}

func (rc *reportController) SummaryHandler(c *gin.Context) {
	report, err := rc.reportService.GetSummaryReport()
	if err != nil {
		log.Printf("Error getting summary report: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching summary report",
		})
		return
	}

	jsonResponse := helper.APIResponse("Summary report fetched successfully", http.StatusOK, "success", report)

	c.JSON(http.StatusOK, jsonResponse)
}

func (rc *reportController) EventHandler(c *gin.Context) {
	eventIDStr := c.DefaultQuery("id", "")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil || eventID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID",
		})
		return
	}

	report, err := rc.reportService.GetEventReport(eventID)
	if err != nil {
		log.Printf("Error getting event report: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching event report",
		})
		return
	}

	jsonResponse := helper.APIResponse("Event report fetched successfully", http.StatusOK, "success", report)

	c.JSON(http.StatusOK, jsonResponse)
}
