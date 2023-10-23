package mock

import (
	"backend/model/entity"
	"context"
	"github.com/stretchr/testify/mock"
	"sync"
)

type CityRepoMock struct {
	mock.Mock
}

// function Provider
func NewCityRepoMock() *CityRepoMock {
	return &CityRepoMock{mock.Mock{}}
}

func (c *CityRepoMock) Insert(ctx context.Context, input *entity.City) (*entity.City, error) {
	args := c.Mock.Called(ctx, input)

	city := args.Get(0)
	if city == nil {
		return nil, args.Error(1)
	}

	return city.(*entity.City), args.Error(1)
}

func (c *CityRepoMock) GetAll(ctx context.Context) ([]entity.City, error) {
	args := c.Mock.Called(ctx)

	cities := args.Get(0)
	if cities == nil {
		return nil, args.Error(1)
	}

	return cities.([]entity.City), args.Error(1)
}

func (c *CityRepoMock) GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.City, chanError chan error) {
	wg.Add(1)
	defer wg.Done()

	args := c.Mock.Called(ctx, wg, id, chanRes, chanError)
	chanRes <- args.Get(0).(entity.City)
	chanError <- args.Error(1)
}

func (c *CityRepoMock) GetByProvinceId(ctx context.Context, wg *sync.WaitGroup, provinceId int, chanRes chan []entity.City, chanError chan error) {
	wg.Add(1)
	defer func() {
		//close(chanError)
		close(chanRes)
		wg.Done()
	}()

	args := c.Mock.Called(ctx, wg, provinceId, chanRes, chanError)
	if cities := args.Get(0); cities == nil {
		chanRes <- nil
	} else {
		chanRes <- cities.([]entity.City)
	}
	chanError <- args.Error(1)
}
