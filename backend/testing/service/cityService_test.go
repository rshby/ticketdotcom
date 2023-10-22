package service_test

import (
	mck "backend/mock"
	"backend/model/dto"
	"backend/model/entity"
	"backend/service"
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTestCity(t *testing.T) {
	getById := dto.CityDetail[*entity.City]{
		Province: &entity.Province{
			Id:   1,
			Name: "DKI Jakarta",
		},
		Cities: &entity.City{
			Id:   1,
			Name: "Jakarta Selatan",
		},
	}

	getByProvinceId := dto.CityDetail[[]entity.City]{
		Province: &entity.Province{
			Id:   1,
			Name: "DKI Jakarta",
		},
		Cities: []entity.City{
			{1, "Jakarta Selatan", 1},
			{2, "Jakarta Pusat", 1},
		},
	}

	fmt.Println(getById)
	fmt.Println(getByProvinceId)
}

func TestCityService_GetAll(t *testing.T) {
	t.Run("get all city success", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock cityRepo method GetAll
		cityRepo.Mock.On("GetAll", mock.Anything).Return([]entity.City{
			{1, "Jakarta Selatan", 1},
			{2, "Bandung", 2},
		}, nil).Times(1)

		// execute service
		result, err := cityService.GetAll(context.Background())

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
		cityRepo.Mock.AssertExpectations(t)
	})
	t.Run("get all city error not found", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock cityRepo method GetAll
		errorMessage := "recor not found"
		cityRepo.Mock.On("GetAll", mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// execute service
		result, err := cityService.GetAll(context.Background())

		// validate testing
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
		cityRepo.Mock.AssertExpectations(t)
	})
}

func TestCityService_GetById(t *testing.T) {
	t.Run("test get city by id success", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock cityRepo method GetById
		cityRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.City{
				Id:         1,
				Name:       "Jakarta Selatan",
				ProvinceId: 1,
			}, nil).Times(1)

		// mock provinceRepo method GetById
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Province{
				Id:   1,
				Name: "DKI Jakarta",
			}, nil).Times(1)

		// execute service
		result, err := cityService.GetById(context.Background(), 1)

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.Province.Id)
		assert.Equal(t, 1, result.Cities.ProvinceId)
		cityRepo.Mock.AssertExpectations(t)
		provinceRepo.Mock.AssertExpectations(t)
	})
	t.Run("test get city by id error city not found", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock cityRepo method GetById
		errorMessage := "record city not found"
		cityRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.City{}, errors.New(errorMessage)).Times(1)

		// execute service method
		result, err := cityService.GetById(context.Background(), 1)

		// validate testing
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Equal(t, errorMessage, err.Error())
		cityRepo.Mock.AssertExpectations(t)
	})
	t.Run("test get city by id error province not found", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock cityRepo method GetById
		cityRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.City{
				Id:         1,
				Name:       "Jakarta Selatan",
				ProvinceId: 1,
			}, nil).Times(1)

		// mock provinceRepo method GetById
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Province{}, errors.New("province not found"))

		// execuet service method
		result, err := cityService.GetById(context.Background(), 1)

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &entity.Province{}, result.Province)
		cityRepo.Mock.AssertExpectations(t)
		provinceRepo.Mock.AssertExpectations(t)
	})
}

func TestCityService_Insert(t *testing.T) {
	t.Run("test insert city success", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock provinceRepo method GetById
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Province{
				Id:   1,
				Name: "DKI Jakarta",
			}, nil).Times(1)

		// mock cityRepo method Insert
		cityName := "Jakarta Selatan"
		cityRepo.Mock.On("Insert", mock.Anything, mock.Anything).Return(&entity.City{
			Id:         1,
			Name:       cityName,
			ProvinceId: 1,
		}, nil).Times(1)

		// execute service method
		result, err := cityService.Insert(context.Background(), &dto.InsertCityRequest{
			Name:       "Jakarta Selatan",
			ProvinceId: 1,
		})

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, result.Id)
		assert.Equal(t, 1, result.ProvinceId)
		assert.Equal(t, cityName, result.Name)
		cityRepo.Mock.AssertExpectations(t)
		provinceRepo.Mock.AssertExpectations(t)
	})
	t.Run("test insert city error province not found", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock provinceRepo method GetById
		errorMessage := "record not found"
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, 1, mock.Anything, mock.Anything).
			Return(entity.Province{}, errors.New(errorMessage)).Times(1)

		// execute service method
		result, err := cityService.Insert(context.Background(), &dto.InsertCityRequest{
			Name:       "Jakarta Selatan",
			ProvinceId: 1,
		})

		// validate testing
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "record province not found", err.Error())
		provinceRepo.Mock.AssertExpectations(t)
	})
	t.Run("test insert city error failed to insert", func(t *testing.T) {
		provinceRepo := mck.NewProvinceRepoMock()
		cityRepo := mck.NewCityRepoMock()
		cityService := service.NewCityService(provinceRepo, cityRepo)

		// mock provinceRepo method GetById
		provinceRepo.Mock.On("GetById", mock.Anything, mock.Anything, 1, mock.Anything, mock.Anything).
			Return(entity.Province{
				Id:   1,
				Name: "DKI Jakarta",
			}, nil).Times(1)

		// mock cityRepo method Insert
		errorMessage := "failed to insert city data"
		cityRepo.Mock.On("Insert", mock.Anything, mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// execute service method
		result, err := cityService.Insert(context.Background(), &dto.InsertCityRequest{
			Name:       "Jakarta Selatan",
			ProvinceId: 1,
		})

		// validate testing
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
		provinceRepo.Mock.AssertExpectations(t)
		cityRepo.Mock.AssertExpectations(t)
	})
}
