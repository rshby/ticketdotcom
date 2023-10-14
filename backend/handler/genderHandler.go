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
	"strings"
)

type GenderHandler struct {
	GenderService service.IGenderService
}

// function provider
func NewGenderHandler(genderService service.IGenderService) *GenderHandler {
	return &GenderHandler{genderService}
}

// handler insert
func (g *GenderHandler) Insert(c *gin.Context) {
	span, ctxTracing := opentracing.StartSpanFromContext(c, "GenderHandler Insert")
	defer span.Finish()

	var request dto.GenderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		statusCode := http.StatusBadRequest
		if msg, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, errMessage := range msg {
				errorMessages = append(errorMessages, fmt.Sprintf("error field %v with tag %v", errMessage.Field(), errMessage.Tag()))
			}

			c.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    strings.Join(errorMessages, ". "),
			})
			return
		}

		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	span.SetTag("request", request)

	res, err := g.GenderService.Insert(ctxTracing, &request)
	if err != nil {
		// if not found
		if strings.Contains(err.Error(), "not found") {
			statusCode := http.StatusNotFound
			c.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    err.Error(),
			})
			return
		}

		// if internal server error
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success insert
	statusCode := http.StatusOK
	c.JSON(statusCode, &dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success insert gender",
		Data:       res,
	})
}

// handler get all data
func (g *GenderHandler) GetAll(c *gin.Context) {
	span, ctxTracing := opentracing.StartSpanFromContext(c, "GenderHander GetAll")
	defer span.Finish()

	res, err := g.GenderService.GetAll(ctxTracing)
	if err != nil {
		// jika error not found
		if strings.Contains(err.Error(), "not found") {
			statusCode := http.StatusNotFound
			c.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    err.Error(),
			})
			return
		}

		// jika error internal server error
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success
	statusCode := http.StatusOK
	c.JSON(statusCode, &dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success get all data gender",
		Data:       res,
	})
}
