package repository

import (
	"context"
	"database/sql"

	"School/internal/domain"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) GetAll(ctx context.Context) ([]domain.Student, error) {
	rows, err := r.db.QueryContext(ctx, `
    SELECT id, name, age, gender, height
    FROM students
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []domain.Student

	for rows.Next() {
		var s domain.Student
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Age,
			&s.Gender,
			&s.Height,
		); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
