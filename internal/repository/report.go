package repository

import (
	"nws/internal/model"

	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) Create(report *model.Reports) error {
	return r.db.Create(report).Error
}

func (r *ReportRepository) GetByID(id int64) (*model.Reports, error) {
	return nil, nil
}

func (r *ReportRepository) GetAll() ([]*model.Reports, error) {
	var reports []*model.Reports
	if err := r.db.Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *ReportRepository) Update(report *model.Reports) error {
	return r.db.Save(report).Error
}

func (r *ReportRepository) Delete(id int64) error {
	return nil
}
