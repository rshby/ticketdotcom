package mock

import (
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type GenderRepoMock struct {
	mock.Mock
}

func NewGenderRepoMock() *GenderRepoMock {
	return &GenderRepoMock{mock.Mock{}}
}

func (g *GenderRepoMock) Insert(ctx context.Context, input *entity.Gender) (*entity.Gender, error) {
	args := g.Mock.Called(ctx, input)

	gender := args.Get(0)
	if gender == nil {
		return nil, args.Error(1)
	}

	return gender.(*entity.Gender), args.Error(1)
}

func (g *GenderRepoMock) GetAll(ctx context.Context) ([]entity.Gender, error) {
	args := g.Mock.Called(ctx)

	gender := args.Get(0)
	if gender == nil {
		return nil, args.Error(1)
	}

	return gender.([]entity.Gender), args.Error(1)
}
