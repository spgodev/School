package controller

import (
	"School/internal/domain"
	"context"
)

type StudentRepo interface {
	Create(ctx context.Context, s domain.Student) (int64, error)
	GetAll(ctx context.Context) ([]domain.Student, error)
}

type ReportBuilder interface {
	BuildReport(ctx context.Context) (*domain.SchoolReport, error)
}
