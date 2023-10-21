package handler_test

import (
	"backend/handler"
	mck "backend/mock"
	"backend/model/dto"
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

func TestProvinceHandler_Insert(t *testing.T) {
	t.Run("test insert province success", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		provinceHandler := handler.NewProvinceHandler(provinceService)

		// mock provinceService method Insert
		provinceService.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Province{
				Id:   1,
				Name: "DKI Jakarta",
			}, nil).Times(1)

		// create routes
		r := gin.Default()
		r.POST("/", provinceHandler.Insert)

		// create request
		request := dto.InsertProvinceRequest{
			Name: "DKI Jakarta",
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

		// validate test
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "ok", responseBody["status"].(string))
		provinceService.Mock.AssertExpectations(t)
	})
	t.Run("test insert province error bad request", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		provinceHandler := handler.NewProvinceHandler(provinceService)

		// create routes
		r := gin.Default()
		r.POST("/", provinceHandler.Insert)

		// create request
		request := dto.InsertProvinceRequest{}
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

		// validate test
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		assert.Equal(t, http.StatusBadRequest, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "bad request", responseBody["status"].(string))
	})
	t.Run("test insert province error internal server error", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		provinceHandler := handler.NewProvinceHandler(provinceService)

		// mock provinceService method Insert
		errorMessage := "failed to insert"
		provinceService.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.POST("/", provinceHandler.Insert)

		// create request
		request := dto.InsertProvinceRequest{
			Name: "DKI Jakarta",
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
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, http.StatusInternalServerError, int(responseBody["status_code"].(float64)))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
		provinceService.Mock.AssertExpectations(t)
	})
}

func TestProvinceHandler_GetById(t *testing.T) {
	t.Run("test get province by id success", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		handler := handler.NewProvinceHandler(provinceService)

		// mock provinceService method GetById
		provinceService.Mock.On("GetById", mock.Anything, mock.Anything).Return(&dto.ProvinceDetail{
			Id:   1,
			Name: "DKI Jakarta",
			Cities: []entity.City{
				{1, "Jakarta Selatan", 1},
				{2, "Jakarta Pusat", 1},
			},
		}, nil).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/:id", handler.GetById)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/1", nil)
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
		provinceService.Mock.AssertExpectations(t)
	})
	t.Run("test get province by id error cant convert id", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		provinceHandler := handler.NewProvinceHandler(provinceService)

		// create routes
		r := gin.Default()
		r.GET("/:id", provinceHandler.GetById)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/asc", nil)
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate testing
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		assert.Equal(t, http.StatusBadRequest, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "bad request", responseBody["status"].(string))
	})
	t.Run("test get province by id error province not found", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		provinceHandler := handler.NewProvinceHandler(provinceService)

		// mock provinceService method GetById
		errorMessage := "record province not found"
		provinceService.Mock.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/:id", provinceHandler.GetById)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/99", nil)
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate testing
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		assert.Equal(t, http.StatusNotFound, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "not found", responseBody["status"].(string))
		provinceService.Mock.AssertExpectations(t)
	})
	t.Run("test get provice by id error internal server error", func(t *testing.T) {
		provinceService := mck.NewProvinceServiceMock()
		provinceHandler := handler.NewProvinceHandler(provinceService)

		// mock provinceService method GetById
		errorMessage := "fail to get province data"
		provinceService.Mock.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/:id", provinceHandler.GetById)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/99", nil)
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		response := recorder.Result()
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate testing
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, http.StatusInternalServerError, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "internal server error", responseBody["status"].(string))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
		provinceService.Mock.AssertExpectations(t)
	})
}
