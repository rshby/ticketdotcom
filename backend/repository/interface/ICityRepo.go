package repository

import (
	"backend/model/entity"
	"context"
	"sync"
)

type ICityRepo interface {
	Insert(ctx context.Context, input *entity.City) (*entity.City, error)
	GetAll(ctx context.Context) ([]entity.City, error)
	GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.City, chanError chan error)
	GetByProvinceId(ctx context.Context, wg *sync.WaitGroup, provinceId int, chanRes chan []entity.City, chanError chan error)
}
