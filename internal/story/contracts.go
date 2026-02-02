package story

import (
	"context"

	"School/internal/domain"
)

type StudentGetter interface {
	GetAll(ctx context.Context) ([]domain.Student, error)
}
