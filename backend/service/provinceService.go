package service

import (
	"backend/model/dto"
	"backend/model/entity"
	repository "backend/repository/interface"
	service "backend/service/interface"
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"sync"
)

type ProvinceService struct {
	ProvinceRepo repository.IProvinceRepo
	CityRepo     repository.ICityRepo
}

// create function provider
func NewProvinceService(provinceRepo repository.IProvinceRepo, cityRepo repository.ICityRepo) service.IProvinceService {
	return &ProvinceService{
		ProvinceRepo: provinceRepo,
		CityRepo:     cityRepo,
	}
}

func (p *ProvinceService) Insert(ctx context.Context, request *dto.InsertProvinceRequest) (*entity.Province, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProvinceService Insert")
	defer span.Finish()

	// create entity
	input := entity.Province{
		Name: request.Name,
	}

	// call procedure insert in repo
	res, err := p.ProvinceRepo.Insert(ctxTracing, &input)
	if err != nil {
		return nil, err
	}

	// success insert
	span.LogFields(
		log.Object("request", *request),
		log.Object("response-province", *res))
	return res, nil
}

func (p *ProvinceService) GetAll(ctx context.Context) ([]entity.Province, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProvinceService GetAll")
	defer span.Finish()

	// call procedure get all in provinceService
	provinces, err := p.ProvinceRepo.GetAll(ctxTracing)
	if err != nil {
		return nil, errors.New("record not found")
	}

	return provinces, nil
}

func (p *ProvinceService) GetById(ctx context.Context, id int) (*dto.ProvinceDetail, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProvinceService GetById")
	defer span.Finish()

	// call procedure
	chanRes := make(chan entity.Province, 1)
	chanErr := make(chan error, 1)
	chanResCities := make(chan []entity.City, 1)
	chanErrCity := make(chan error, 1)

	wg := &sync.WaitGroup{}
	go p.ProvinceRepo.GetById(ctxTracing, wg, id, chanRes, chanErr)               // get province
	go p.CityRepo.GetByProvinceId(ctxTracing, wg, id, chanResCities, chanErrCity) // get cities by province_id
	wg.Wait()

	// receive from channel
	province, err := <-chanRes, <-chanErr
	if err != nil {
		return nil, err
	}

	cities, _ := <-chanResCities, <-chanErrCity

	// success
	response := dto.ProvinceDetail{
		Id:     province.Id,
		Name:   province.Name,
		Cities: cities,
	}
	return &response, nil
}
