package mock

import (
	"backend/model/dto"
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type ProvinceServiceMock struct {
	mock.Mock
}

func NewProvinceServiceMock() *ProvinceServiceMock {
	return &ProvinceServiceMock{mock.Mock{}}
}

func (p *ProvinceServiceMock) Insert(ctx context.Context, request *dto.InsertProvinceRequest) (*entity.Province, error) {
	args := p.Mock.Called(ctx, request)

	province := args.Get(0)
	if province == nil {
		return nil, args.Error(1)
	}

	return province.(*entity.Province), args.Error(1)
}

func (p *ProvinceServiceMock) GetAll(ctx context.Context) ([]entity.Province, error) {
	args := p.Mock.Called(ctx)

	provinces := args.Get(0)
	if provinces == nil {
		return nil, args.Error(1)
	}

	return provinces.([]entity.Province), args.Error(1)
}

func (p *ProvinceServiceMock) GetById(ctx context.Context, id int) (*dto.ProvinceDetail, error) {
	args := p.Mock.Called(ctx, id)

	province := args.Get(0)
	if province == nil {
		return nil, args.Error(1)
	}

	return province.(*dto.ProvinceDetail), args.Error(1)
}
