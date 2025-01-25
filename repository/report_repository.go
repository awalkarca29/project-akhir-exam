package repository

import (
	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSummaryReport() (map[string]interface{}, error)
	GetEventReport(eventID int) (map[string]interface{}, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetSummaryReport() (map[string]interface{}, error) {
	var result map[string]interface{}

	query := `
        SELECT
            SUM(tickets.quantity) AS total_tickets_sold,
            SUM(tickets.total) AS total_revenue
        FROM tickets
    `
	err := r.db.Raw(query).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *reportRepository) GetEventReport(eventID int) (map[string]interface{}, error) {
	var result map[string]interface{}

	query := `
        SELECT
            events.name AS event_name,
            SUM(tickets.quantity) AS total_tickets_sold,
            SUM(tickets.total) AS total_revenue
        FROM tickets
        JOIN events ON tickets.event_id = events.id
        WHERE tickets.event_id = ?
        GROUP BY events.name
    `
	err := r.db.Raw(query, eventID).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
