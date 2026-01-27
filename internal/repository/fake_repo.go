package repository

import (
	"School/internal/domain"
	"context"
)

type FakeRepo struct {
	Students []domain.Student
	Err      error
}

func (f FakeRepo) GetAll(ctx context.Context) ([]domain.Student, error) {
	return f.Students, f.Err
}
