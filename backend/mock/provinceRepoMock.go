package mock

import (
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
	"sync"
)

type ProvinceRepoMock struct {
	mock.Mock
}

func (p *ProvinceRepoMock) Insert(ctx context.Context, input *entity.Province) (*entity.Province, error) {
	args := p.Mock.Called(ctx, input)

	province := args.Get(0)
	if province == nil {
		return nil, args.Error(1)
	}

	return province.(*entity.Province), args.Error(1)
}

func (p *ProvinceRepoMock) GetAll(ctx context.Context) ([]entity.Province, error) {
	args := p.Mock.Called(ctx)

	provinces := args.Get(0)
	if provinces == nil {
		return nil, args.Error(1)
	}

	return provinces.([]entity.Province), args.Error(1)
}

func (p *ProvinceRepoMock) GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.Province, chanError chan error) {
	wg.Add(1)
	defer func() {
		close(chanError)
		close(chanRes)
		wg.Done()
	}()

	args := p.Mock.Called(ctx, wg, id, chanRes, chanError)
	chanRes <- args.Get(0).(entity.Province)
	chanError <- args.Error(1)
}

// function provider
func NewProvinceRepoMock() *ProvinceRepoMock {
	return &ProvinceRepoMock{mock.Mock{}}
}
