package repository

import (
	"backend/model/entity"
	"context"
	"sync"
)

type IProvinceRepo interface {
	Insert(ctx context.Context, input *entity.Province) (*entity.Province, error)
	GetAll(ctx context.Context) ([]entity.Province, error)
	GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.Province, chanError chan error)
}
