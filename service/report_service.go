package service

import (
	"log"
	"project-akhir-exam/repository"
)

type ReportService interface {
	GetSummaryReport() (map[string]interface{}, error)
	GetEventReport(eventID int) (map[string]interface{}, error)
}

type reportService struct {
	reportRepository repository.ReportRepository
}

func NewReportService(reportRepository repository.ReportRepository) ReportService {
	return &reportService{reportRepository: reportRepository}
}

func (s *reportService) GetSummaryReport() (map[string]interface{}, error) {
	report, err := s.reportRepository.GetSummaryReport()
	if err != nil {
		log.Println("Error getting summary report:", err)
		return nil, err
	}
	return report, nil
}

func (s *reportService) GetEventReport(eventID int) (map[string]interface{}, error) {
	report, err := s.reportRepository.GetEventReport(eventID)
	if err != nil {
		log.Println("Error getting event report:", err)
		return nil, err
	}
	return report, nil
}
