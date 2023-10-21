package service_test

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

func TestGenderService_Insert(t *testing.T) {
	t.Run("test insert gender success", func(t *testing.T) {
		genderRepo := mck.NewGenderRepoMock()
		genderService := service.NewGenderService(genderRepo)

		// mock genderRepo method Insert
		genderRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Gender{1, "M", "Male"}, nil).Times(1)

		// execute service method
		result, err := genderService.Insert(context.Background(), &dto.GenderRequest{
			Code: "M",
			Name: "Male",
		})

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "M", result.Code)
		assert.Equal(t, "Male", result.Name)
		genderRepo.Mock.AssertExpectations(t)
	})
	t.Run("test insert gender error", func(t *testing.T) {
		genderRepo := mck.NewGenderRepoMock()
		genderService := service.NewGenderService(genderRepo)

		// mock genderRepo method Insert
		errorMessage := "record not found"
		genderRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, errors.New(errorMessage)).Times(1)

		// execute service method
		result, err := genderService.Insert(context.Background(), &dto.GenderRequest{
			Code: "F",
			Name: "Female",
		})

		// validate testing
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
		genderRepo.Mock.AssertExpectations(t)
	})
}

func TestGenderService_GetAll(t *testing.T) {
	t.Run("test get all gender succes", func(t *testing.T) {
		genderRepo := mck.NewGenderRepoMock()
		genderService := service.NewGenderService(genderRepo)

		// mock genderRepo method GetAll
		genderRepo.Mock.On("GetAll", mock.Anything).
			Return([]entity.Gender{
				{1, "M", "Male"},
				{2, "F", "Female"},
			}, nil).Times(1)

		// execuet service method
		result, err := genderService.GetAll(context.Background())

		// validate testing
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, len(result))
		genderRepo.Mock.AssertExpectations(t)
	})
	t.Run("test get all gender error not found", func(t *testing.T) {
		genderRepo := mck.NewGenderRepoMock()
		genderService := service.NewGenderService(genderRepo)

		// mock genderRepo method GetAll
		genderRepo.Mock.On("GetAll", mock.Anything).Return(nil, errors.New("record not found")).Times(1)

		// execute service method
		result, err := genderService.GetAll(context.Background())

		// validate testing
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "record not found", err.Error())
		genderRepo.Mock.AssertExpectations(t)
	})
}
