package handler

import (
	"backend/helper"
	"backend/model/dto"
	service "backend/service/interface"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strconv"
	"strings"
)

type CityHandler struct {
	CityService service.ICityService
}

// function provider
func NewCityHandler(cityService service.ICityService) *CityHandler {
	return &CityHandler{
		CityService: cityService,
	}
}

// handler get all city
func (c *CityHandler) GetAllCity(ctx *gin.Context) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityHandler GetAllCity")
	defer span.Finish()

	// call procedure get all in service
	cities, err := c.CityService.GetAll(ctxTracing)
	if err != nil {
		statusCode := http.StatusNotFound
		ctx.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success get all cities
	statusCode := http.StatusOK
	ctx.JSON(statusCode, &dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success get all data cities",
		Data:       cities,
	})
	return
}

// handler insert city
func (c *CityHandler) Insert(ctx *gin.Context) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityHandler Insert")
	defer span.Finish()

	var request dto.InsertCityRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if errMsg, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, item := range errMsg {
				errorMessages = append(errorMessages, fmt.Sprintf("error field %v, with tag %v", item.Field(), item.Tag()))
			}

			statusCode := http.StatusBadRequest
			ctx.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    strings.Join(errorMessages, ". "),
			})
			return
		}
	}

	// call procedure insert in service
	result, err := c.CityService.Insert(ctxTracing, &request)
	if err != nil {
		// if error not found
		if strings.Contains(err.Error(), "not found") {
			statusCode := http.StatusNotFound
			ctx.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    err.Error(),
			})
			return
		}

		// if error internal server
		statusCode := http.StatusInternalServerError
		ctx.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success insert
	statusCode := http.StatusOK
	response := dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success insert city to database",
		Data:       result,
	}

	span.SetTag("response-object", response)
	ctx.JSON(statusCode, &response)
	return
}

// handler get city by province_id
func (c *CityHandler) GetByProvinceId(ctx *gin.Context) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityHandler GetByProvinceId")
	defer span.Finish()

	// get id
	provinceId, err := strconv.Atoi(ctx.Params.ByName("provinceId"))
	if err != nil {
		statusCode := http.StatusBadRequest
		ctx.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    "cant convert id to int",
		})
		return
	}

	// call procedure in service
	cities, err := c.CityService.GetByProvinceId(ctxTracing, provinceId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			statusCode := http.StatusNotFound
			ctx.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    err.Error(),
			})
			return
		}

		// if internal server error
		statusCode := http.StatusInternalServerError
		ctx.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success
	statusCode := http.StatusOK
	ctx.JSON(statusCode, &dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success get data cities by province_id",
		Data:       cities,
	})
	return
}
