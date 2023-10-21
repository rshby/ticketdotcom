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
	"strconv"
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

// handler get province by id
func (p *ProvinceHandler) GetById(c *gin.Context) {
	span, ctxTracing := opentracing.StartSpanFromContext(c, "ProvinceHandler GetById")
	defer span.Finish()

	// get id
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(http.StatusBadRequest),
			Message:    "cant convert id to int",
		})
		return
	}

	span.SetTag("request-id", id)

	// call method GetById in service
	result, err := p.ProvinceService.GetById(ctxTracing, id)
	if err != nil {
		// jika not found
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no rows") {
			statusCode := http.StatusNotFound
			c.JSON(statusCode, &dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.GenerateStatusFromCode(statusCode),
				Message:    "record province not found",
			})
			return
		}

		// jika internal server error
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    err.Error(),
		})
		return
	}

	// success get data province by id
	span.SetTag("response-object", *result)
	statusCode := http.StatusOK
	c.JSON(statusCode, &dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.GenerateStatusFromCode(statusCode),
		Message:    "success get data province by id",
		Data:       result,
	})
}
