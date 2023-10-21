package testing

import (
	mck "backend/mock"
	"backend/model/dto"
	"backend/model/entity"
	"backend/service"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProvinceService_Insert(t *testing.T) {
	t.Run("testing insert province success", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		provinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method Insert
		provinceRepo.Mock.On("Insert", mock.Anything, mock.Anything).Return(&entity.Province{
			Id:   1,
			Name: "DKI Jakarta",
		}, nil).Times(1)

		// execuet service method
		res, err := provinceService.Insert(context.Background(), &dto.InsertProvinceRequest{
			Name: "DKI Jakarta",
		})

		// validate testing result
		assert.Nil(t, err)
		assert.NotNil(t, res)
		provinceRepo.Mock.AssertExpectations(t)
	})
	t.Run("testing insert province error", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		provinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method Insert
		errorMessage := "record not found"
		provinceRepo.Mock.On("Insert", mock.Anything, mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// execute service method
		res, err := provinceService.Insert(context.Background(), &dto.InsertProvinceRequest{Name: "DKI Jakarta"})

		// validate testing
		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.Equal(t, errorMessage, err.Error())
		provinceRepo.Mock.AssertExpectations(t)
	})
}

func TestProvinceService_GetAll(t *testing.T) {
	t.Run("test get all province success", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		proviinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method GetAll
		provinceRepo.Mock.On("GetAll", mock.Anything).Return([]entity.Province{
			{1, "DKI Jakarta"},
			{2, "Jawa Barat"},
		}, nil).Times(1)

		// execuet service method
		result, err := proviinceService.GetAll(context.Background())

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
		provinceRepo.Mock.AssertExpectations(t)
	})
	t.Run("test get all province error not found", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		provinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method GetAll
		provinceRepo.Mock.On("GetAll", mock.Anything).Return(nil, errors.New("not found")).Times(1)

		// execute service method
		result, err := provinceService.GetAll(context.Background())

		// validate testing
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "record not found", err.Error())
		provinceRepo.Mock.AssertExpectations(t)
	})
}

func TestProvinceService_GetById(t *testing.T) {
	t.Run("test get province by id success", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		provinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(entity.Province{
			Id:   1,
			Name: "DKI Jakarta",
		}, nil).Times(1)

		// mock cityRepo method GetByProvinceId
		cityRepo.Mock.On("GetByProvinceId", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]entity.City{
			{1, "Jakarta Selatan", 1},
			{2, "Jakarta Pusat", 1},
		}, nil).Times(1)

		// execute service method
		result, err := provinceService.GetById(context.Background(), 1)

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result.Cities))
		provinceRepo.Mock.AssertExpectations(t)
		cityRepo.Mock.AssertExpectations(t)
	})
	t.Run("test get province by id error province not found", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		provinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method GetById
		errorMessage := "record province not found"
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(entity.Province{}, errors.New(errorMessage)).Times(1)

		// mock cityRepo method GetByProvinceId
		cityRepo.Mock.On("GetByProvinceId", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("record cities not found")).Times(1)

		// execute service method
		result, err := provinceService.GetById(context.Background(), 1)

		// validate testing
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, errorMessage, err.Error())
		provinceRepo.Mock.AssertExpectations(t)
		cityRepo.Mock.AssertExpectations(t)
	})
	t.Run("test get province by id success but cities not exist", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		provinceService := service.NewProvinceService(provinceRepo, cityRepo)

		// mock provinceRepo method GetById
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Province{
				Id:   1,
				Name: "DKI Jakarta",
			}, nil).Times(1)

		// mock cityRepo method GetByProvinceId
		cityRepo.Mock.On("GetByProvinceId", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, errors.New("cities not found")).Times(1)

		// execuete service method
		result, err := provinceService.GetById(context.Background(), 1)

		// validate testing
		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Nil(t, result.Cities)
		assert.Equal(t, 1, result.Id)
		provinceRepo.Mock.AssertExpectations(t)
		cityRepo.Mock.AssertExpectations(t)
	})
}
