package mock

import (
	"backend/model/dto"
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type GenderServiceMock struct {
	mock.Mock
}

// function provider
func NewGenderServiceMock() *GenderServiceMock {
	return &GenderServiceMock{mock.Mock{}}
}

func (g *GenderServiceMock) Insert(ctx context.Context, request *dto.GenderRequest) (*entity.Gender, error) {
	args := g.Mock.Called(ctx, request)

	gender := args.Get(0)
	if gender == nil {
		return nil, args.Error(1)
	}

	return gender.(*entity.Gender), args.Error(1)
}

func (g *GenderServiceMock) GetAll(ctx context.Context) ([]entity.Gender, error) {
	args := g.Mock.Called(ctx)

	gender := args.Get(0)
	if gender == nil {
		return nil, args.Error(1)
	}

	return gender.([]entity.Gender), args.Error(1)
}
