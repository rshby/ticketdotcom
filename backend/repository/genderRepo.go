package repository

import (
	"backend/model/entity"
	"backend/repository/interface"
	"context"
	"database/sql"
	"errors"
	"github.com/opentracing/opentracing-go"
)

type GenderRepo struct {
	DB *sql.DB
}

// function provider
func NewGenderRepo(db *sql.DB) repository.IGenderRepo {
	return &GenderRepo{db}
}

func (g *GenderRepo) Insert(ctx context.Context, input *entity.Gender) (*entity.Gender, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "GenderRepo Insert")
	defer span.Finish()

	statement, err := g.DB.PrepareContext(ctxTracing, "INSERT INTO gender(code, name) VALUES (?, ?)")
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	result, err := statement.ExecContext(ctxTracing, input.Code, input.Name)
	if err != nil {
		return nil, err
	}

	if row, _ := result.RowsAffected(); row == 0 {
		return nil, errors.New("error failed to insert data gender")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	input.Id = int(id)
	return input, nil
}

func (g *GenderRepo) GetAll(ctx context.Context) ([]entity.Gender, error) {
	// create context tracing
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "GenderRepo GetAll")
	defer span.Finish()

	statement, err := g.DB.PrepareContext(ctxTracing, "SELECT id, code, name FROM gender")
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	// execute
	rows, err := statement.QueryContext(ctxTracing)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var genders []entity.Gender
	for rows.Next() {
		var gender entity.Gender
		if err := rows.Scan(&gender.Id, &gender.Code, &gender.Name); err != nil {
			return nil, err
		}

		// append to object response
		genders = append(genders, gender)
	}

	// if not found
	if len(genders) == 0 {
		return nil, errors.New("record not found")
	}

	// success
	return genders, nil
}
