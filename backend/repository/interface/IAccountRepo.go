package repository

import (
	"backend/model/entity"
	"context"
)

type IAccountRepo interface {
	Insert(ctx context.Context, input *entity.Account) (*entity.Account, error)
	GetAll(ctx context.Context) ([]entity.Account, error)
}
