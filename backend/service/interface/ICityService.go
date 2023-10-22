package service

import (
	"backend/model/dto"
	"backend/model/entity"
	"context"
)

type ICityService interface {
	GetAll(ctx context.Context) ([]entity.City, error)
	GetById(ctx context.Context, id int) (*dto.CityDetail[*entity.City], error)
	GetByProvinceId(ctx context.Context, provinceId int) (*dto.CityDetail[[]entity.City], error)
}
