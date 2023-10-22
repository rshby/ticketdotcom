package handler_test

import (
	"backend/handler"
	mck "backend/mock"
	"backend/model/entity"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCityHandler_GetAll(t *testing.T) {
	t.Run("test get all city success", func(t *testing.T) {
		cityService := mck.NewCityServiceMock()
		cityHandler := handler.NewCityHandler(cityService)

		// mock cityService method GetAll
		cityService.Mock.On("GetAll", mock.Anything).Return([]entity.City{
			{1, "Jakarta Selatan", 1},
		}, nil).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/", cityHandler.GetAllCity)

		// create req
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate test
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "ok", responseBody["status"].(string))
		assert.Equal(t, "success get all data cities", responseBody["message"].(string))
		cityService.Mock.AssertExpectations(t)
	})
	t.Run("test get all city error not found", func(t *testing.T) {
		cityService := mck.NewCityServiceMock()
		cityHandler := handler.NewCityHandler(cityService)

		// mock cityService method GetAll
		errorMessage := "record cities not found"
		cityService.Mock.On("GetAll", mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/", cityHandler.GetAllCity)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate test
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		assert.Equal(t, http.StatusNotFound, int(responseBody["status_code"].(float64)))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
		assert.Equal(t, "not found", responseBody["status"].(string))
		cityService.Mock.AssertExpectations(t)
	})
}

func TestCityHandler_Insert(t *testing.T) {
	t.Run("test insert city success", func(t *testing.T) {
		cityService := mck.NewCityServiceMock()
		cityHandler := handler.NewCityHandler(cityService)

		// mock cityService method Insert
		cityService.Mock.On("Insert", mock.Anything, mock.Anything).Return(&entity.City{
			Id:         1,
			Name:       "Jakarta Selatan",
			ProvinceId: 1,
		}, nil).Times(1)

		// create routes
		r := gin.Default()
		r.POST("/", cityHandler.Insert)

		// create request
		request := map[string]any{
			"name":        "Jakarta Selatan",
			"province_id": 1,
		}
		reqJson, _ := json.Marshal(&request)
		requestBody := strings.NewReader(string(reqJson))

		req := httptest.NewRequest(http.MethodPost, "/", requestBody)
		req.Header.Add("content-type", "application/json")
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate testing
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "ok", responseBody["status"].(string))
		cityService.Mock.AssertExpectations(t)
	})
}
