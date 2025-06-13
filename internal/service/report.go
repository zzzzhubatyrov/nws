package service

import (
	"nws/internal/model"
	"nws/internal/repository"
)

type ReportService struct {
	repo repository.ReportRepositoryInterface
}

func NewReportService(repo repository.ReportRepositoryInterface) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) CreateReport(report *model.Reports) error {
	return s.repo.Create(report)
}

func (s *ReportService) GetReport(id int64) (*model.Reports, error) {
	return s.repo.GetByID(id)
}

func (s *ReportService) GetAllReports() ([]*model.Reports, error) {
	reports, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (s *ReportService) UpdateReport(report *model.Reports) error {
	return s.repo.Update(report)
}

func (s *ReportService) DeleteReport(id int64) error {
	return s.repo.Delete(id)
}
