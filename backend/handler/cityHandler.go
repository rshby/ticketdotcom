package handler

import (
	"backend/helper"
	"backend/model/dto"
	service "backend/service/interface"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"net/http"
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

// method get all city
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
