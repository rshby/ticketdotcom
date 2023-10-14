package service

import (
	"backend/model/dto"
	"backend/model/entity"
	repository "backend/repository/interface"
	service "backend/service/interface"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
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
	//TODO implement me
	panic("implement me")
}

func (p *ProvinceService) GetById(ctx context.Context, id int) (*dto.ProvinceDetail, error) {
	//TODO implement me
	panic("implement me")
}
