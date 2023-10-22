package service

import (
	"backend/model/dto"
	"backend/model/entity"
	repository "backend/repository/interface"
	service "backend/service/interface"
	"context"
	"github.com/opentracing/opentracing-go"
	"sync"
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

func (c *CityService) GetById(ctx context.Context, id int) (*dto.CityDetail[*entity.City], error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityService GetById")
	defer span.Finish()

	// call cityRepo method get by id
	wg := &sync.WaitGroup{}
	chanCity := make(chan entity.City, 1)
	chanErr := make(chan error, 1)
	go c.CityRepository.GetById(ctxTracing, wg, id, chanCity, chanErr)
	wg.Wait()

	city, err := <-chanCity, <-chanErr
	if err != nil {
		return nil, err
	}

	// get data province
	chanProvince := make(chan entity.Province, 1)
	go c.ProvinceRepository.GetById(ctxTracing, wg, city.ProvinceId, chanProvince, chanErr)
	wg.Wait()

	province, _ := <-chanProvince, <-chanErr

	// success
	response := &dto.CityDetail[*entity.City]{
		Province: &province,
		Cities:   &city,
	}
	span.SetTag("response-object", *response)
	return response, nil
}

func (c *CityService) GetByProvinceId(ctx context.Context, provinceId int) (*dto.CityDetail[[]entity.City], error) {
	//TODO implement me
	panic("implement me")
}
