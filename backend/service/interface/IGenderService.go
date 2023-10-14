package service

import (
	"backend/model/dto"
	"backend/model/entity"
	"context"
)

type IGenderService interface {
	Insert(ctx context.Context, request *dto.GenderRequest) (*entity.Gender, error)
	GetAll(ctx context.Context) ([]entity.Gender, error)
}
