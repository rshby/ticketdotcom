package service

import (
	"backend/model/dto"
	"backend/model/entity"
	"context"
)

type IProvinceService interface {
	Insert(ctx context.Context, request *dto.InsertProvinceRequest) (*entity.Province, error)
	GetAll(ctx context.Context) ([]entity.Province, error)
	GetById(ctx context.Context, id int) (*dto.ProvinceDetail, error)
}
