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

	statement, err := g.DB.PrepareContext(ctxTracing, "INSERT INTO gender(code, name) VALUES ($1, $2)")
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
	//TODO implement me
	panic("implement me")
}
