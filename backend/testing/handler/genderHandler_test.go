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

func TestGenderHandler_Insert(t *testing.T) {
	t.Run("test insert gender success", func(t *testing.T) {
		genderService := mck.NewGenderServiceMock()
		genderHandler := handler.NewGenderHandler(genderService)

		// mock genderService method Insert
		genderService.Mock.On("Insert", mock.Anything, mock.Anything).Return(&entity.Gender{
			Id:   1,
			Code: "M",
			Name: "Male",
		}, nil).Times(1)

		// create routes
		r := gin.Default()
		r.POST("/", genderHandler.Insert)

		// create request
		request := dto.GenderRequest{
			Code: "M",
			Name: "Male",
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
		genderService.Mock.AssertExpectations(t)
	})
	t.Run("test insert gender error bad request", func(t *testing.T) {
		genderService := mck.NewGenderServiceMock()
		genderHandler := handler.NewGenderHandler(genderService)

		// create routes
		r := gin.Default()
		r.POST("/", genderHandler.Insert)

		// create request
		request := dto.GenderRequest{
			Code: "MM",
			Name: "Male",
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
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		assert.Equal(t, http.StatusBadRequest, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "bad request", responseBody["status"].(string))
	})
	t.Run("test insert gender error internal server error", func(t *testing.T) {
		genderService := mck.NewGenderServiceMock()
		genderHandler := handler.NewGenderHandler(genderService)

		// mock genderService method Insert
		errorMessage := "failed to insert data gender"
		genderService.Mock.On("Insert", mock.Anything, mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.POST("/", genderHandler.Insert)

		// create request
		request := dto.GenderRequest{
			Code: "M",
			Name: "Male",
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
		genderService.Mock.AssertExpectations(t)
	})
}

func TestGenderHandler_GetAll(t *testing.T) {
	t.Run("get all gender success", func(t *testing.T) {
		genderService := mck.NewGenderServiceMock()
		genderHandler := handler.NewGenderHandler(genderService)

		// mock genderService method GetAll
		genderService.Mock.On("GetAll", mock.Anything).Return([]entity.Gender{
			{1, "F", "Female"},
			{2, "M", "Male"},
		}, nil)

		// create routes
		r := gin.Default()
		r.GET("/", genderHandler.GetAll)

		// create request
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		recorder := httptest.NewRecorder()

		// execute handler
		r.ServeHTTP(recorder, req)

		// result
		result := recorder.Result()
		body, _ := io.ReadAll(result.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		// validate test
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "ok", responseBody["status"].(string))
	})
	t.Run("get all gender error not found", func(t *testing.T) {
		genderService := mck.NewGenderServiceMock()
		genderHandler := handler.NewGenderHandler(genderService)

		// mock genderService method GetAll
		errorMessage := "record not found"
		genderService.Mock.On("GetAll", mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/", genderHandler.GetAll)

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

		// validate testing
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		assert.Equal(t, http.StatusNotFound, int(responseBody["status_code"].(float64)))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
		genderService.Mock.AssertExpectations(t)
	})
	t.Run("get all gender error internal server error", func(t *testing.T) {
		genderService := mck.NewGenderServiceMock()
		genderHandler := handler.NewGenderHandler(genderService)

		// mock genderService method GetAll
		errorMessage := "database error"
		genderService.Mock.On("GetAll", mock.Anything).Return(nil, errors.New(errorMessage)).Times(1)

		// create routes
		r := gin.Default()
		r.GET("/", genderHandler.GetAll)

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

		// validate testing
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
		assert.Equal(t, http.StatusInternalServerError, int(responseBody["status_code"].(float64)))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
		genderService.Mock.AssertExpectations(t)
	})
}
