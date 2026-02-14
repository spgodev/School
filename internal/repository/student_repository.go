package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"School/internal/domain"
)

type StudentRepository struct {
	db *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) GetAll(ctx context.Context) ([]domain.Student, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, age, gender, height
		FROM students
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := make([]domain.Student, 0)

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

func (r *StudentRepository) Create(ctx context.Context, s domain.Student) (int64, error) {
	var id int64
	err := r.db.QueryRow(ctx, `
		INSERT INTO students (name, age, gender, height)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, s.Name, s.Age, s.Gender, s.Height).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
