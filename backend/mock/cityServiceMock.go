package mock

import (
	"backend/model/dto"
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type CityServiceMock struct {
	mock.Mock
}

func NewCityServiceMock() *CityServiceMock {
	return &CityServiceMock{
		mock.Mock{},
	}
}

func (c *CityServiceMock) Insert(ctx context.Context, request *dto.InsertCityRequest) (*entity.City, error) {
	args := c.Mock.Called(ctx, request)

	city := args.Get(0)
	if city == nil {
		return nil, args.Error(1)
	}

	return city.(*entity.City), args.Error(1)
}

func (c *CityServiceMock) GetAll(ctx context.Context) ([]entity.City, error) {
	args := c.Mock.Called(ctx)

	cities := args.Get(0)
	if cities == nil {
		return nil, args.Error(1)
	}

	return cities.([]entity.City), args.Error(1)
}

func (c *CityServiceMock) GetById(ctx context.Context, id int) (*dto.CityDetail[*entity.City], error) {
	args := c.Mock.Called(ctx, id)

	city := args.Get(0)
	err := args.Error(1)
	if city == nil {
		return nil, err
	}

	return city.(*dto.CityDetail[*entity.City]), err
}

func (c *CityServiceMock) GetByProvinceId(ctx context.Context, provinceId int) (*dto.CityDetail[[]entity.City], error) {
	args := c.Mock.Called(ctx, provinceId)

	city := args.Get(0)
	err := args.Error(1)
	if city == nil {
		return nil, err
	}

	return city.(*dto.CityDetail[[]entity.City]), err
}
