package service

import (
	"backend/model/dto"
	"backend/model/entity"
	repository "backend/repository/interface"
	service "backend/service/interface"
	"context"
	"github.com/opentracing/opentracing-go"
)

type CityService struct {
	ProvinceRepository repository.IProvinceRepo
	CityRepository     repository.ICityRepo
}

// function provider
func NewCityService(provinceRepo repository.IProvinceRepo, cityRepo repository.ICityRepo) service.ICityService {
	return &CityService{
		ProvinceRepository: provinceRepo,
		CityRepository:     cityRepo,
	}
}

func (c *CityService) GetAll(ctx context.Context) ([]entity.City, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityService GetAll")
	defer span.Finish()

	// call procedure in repository
	cities, err := c.CityRepository.GetAll(ctxTracing)
	if err != nil {
		return nil, err
	}

	// success get all data
	return cities, nil
}

func (c *CityService) GetById(ctx context.Context) (*dto.CityDetail[*entity.City], error) {
	//TODO implement me
	panic("implement me")
}

func (c *CityService) GetByProvinceId(ctx context.Context, provinceId int) (*dto.CityDetail[[]entity.City], error) {
	//TODO implement me
	panic("implement me")
}
