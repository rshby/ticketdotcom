package service

import (
	"backend/model/dto"
	"backend/model/entity"
	repository "backend/repository/interface"
	service "backend/service/interface"
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
)

type GenderService struct {
	GenderRepo repository.IGenderRepo
}

func (g *GenderService) Insert(ctx context.Context, request *dto.GenderRequest) (*entity.Gender, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "GenderService Insert")
	defer span.Finish()

	span.SetTag("request", request)

	// create entity
	input := entity.Gender{
		Code: request.Code,
		Name: request.Name,
	}

	// insert
	res, err := g.GenderRepo.Insert(ctxTracing, &input)
	if err != nil {
		return nil, err
	}

	// success insert
	return res, nil
}

func (g *GenderService) GetAll(ctx context.Context) ([]entity.Gender, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "GenderService GetAll")
	defer span.Finish()

	// call get all method in repository
	res, err := g.GenderRepo.GetAll(ctxTracing)
	if err != nil {
		return nil, errors.New("record not found")
	}

	// success
	return res, nil
}

// function provider
func NewGenderService(genderRepo repository.IGenderRepo) service.IGenderService {
	return &GenderService{
		GenderRepo: genderRepo,
	}
}
