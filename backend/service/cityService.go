package service

import (
	"backend/model/dto"
	"backend/model/entity"
	repository "backend/repository/interface"
	service "backend/service/interface"
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"reflect"
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

func (c *CityService) Insert(ctx context.Context, request *dto.InsertCityRequest) (*entity.City, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityService Insert")
	defer span.Finish()

	// cek province apakah ada di database
	wg := &sync.WaitGroup{}
	chanProvince := make(chan entity.Province, 1)
	chanErr := make(chan error, 1)
	go c.ProvinceRepository.GetById(ctxTracing, wg, request.ProvinceId, chanProvince, chanErr)
	wg.Wait()
	_, err := <-chanProvince, <-chanErr
	if err != nil {
		return nil, errors.New("record province not found")
	}

	// create entity
	entity := entity.City{
		Name:       request.Name,
		ProvinceId: request.ProvinceId,
	}

	// call procedure insert in cityRepo
	result, err := c.CityRepository.Insert(ctxTracing, &entity)
	if err != nil {
		return nil, err
	}

	// success insert
	return result, nil
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
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityService GetProvinceById")
	defer span.Finish()

	span.SetTag("request-province-id", provinceId)

	// get province with goroutine
	wg := &sync.WaitGroup{}
	chanProvince := make(chan entity.Province, 1)
	chanCity := make(chan []entity.City, 1)
	chanErr := make(chan error, 2)
	defer close(chanErr)

	go c.ProvinceRepository.GetById(ctxTracing, wg, provinceId, chanProvince, chanErr)
	go c.CityRepository.GetByProvinceId(ctxTracing, wg, provinceId, chanCity, chanErr)

	wg.Wait()

	province, _ := <-chanProvince, <-chanErr
	city := <-chanCity

	span.SetTag("response-province", province)
	span.SetTag("response-city", city)

	// if province not found
	if reflect.DeepEqual(province, entity.Province{}) {
		return nil, errors.New("record province not found")
	}

	// success get data city by province id
	response := dto.CityDetail[[]entity.City]{
		Province: &province,
		Cities:   city,
	}

	span.SetTag("response-object", response)
	return &response, nil
}
