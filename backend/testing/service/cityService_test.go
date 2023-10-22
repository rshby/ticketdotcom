package service_test

import (
	mck "backend/mock"
	"backend/model/dto"
	"backend/model/entity"
	"backend/service"
	"context"
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
}
