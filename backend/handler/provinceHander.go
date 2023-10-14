package handler

import (
	"backend/helper"
	"backend/model/dto"
	service "backend/service/interface"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"strings"
)

type ProvinceHandler struct {
	ProvinceService service.IProvinceService
}

// function provider
func NewProvinceHandler(provinceService service.IProvinceService) *ProvinceHandler {
	return &ProvinceHandler{provinceService}
}

// handler insert province
func (p *ProvinceHandler) Insert(c *gin.Context) {
	// create span tracer
	span, ctxTracing := opentracing.StartSpanFromContext(c, "ProvinceHandler Insert")
	defer span.Finish()

	var request dto.InsertProvinceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		statusCode := http.StatusBadRequest
		if msg, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, errMsg := range msg {
				errorMessages = append(errorMessages, fmt.Sprintf("error field %v with tag %v", errMsg.Field(), errMsg.Tag()))
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

	// call procedure insert in repository
	res, err := p.ProvinceService.Insert(ctxTracing, &request)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success
	span.LogFields(
		log.Object("request-body", request),
		log.Object("response-province", *res),
	)

	statusCode := http.StatusOK
	c.JSON(statusCode, &dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success insert province",
		Data:       res,
	})
	return
}
