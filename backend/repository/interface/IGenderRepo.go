package repository

import (
	"backend/model/entity"
	"context"
)

type IGenderRepo interface {
	Insert(ctx context.Context, input *entity.Gender) (*entity.Gender, error)
	GetAll(ctx context.Context) ([]entity.Gender, error)
}
