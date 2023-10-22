package main

import (
	"backend/database"
	"backend/handler"
	"backend/helper"
	"backend/model/dto"
	"backend/repository"
	"backend/router"
	"backend/service"
	"backend/tracing"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("cant load env file : ", err.Error())
	}
}

func main() {
	fmt.Println("app run")

	db := database.ConnectDB()
	fmt.Println(db)

	// register tracing
	tracer, closer, err := tracing.ConnectJaegerTracing()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		log.Fatalf("err cant connect jaeger : " + err.Error())
	}

	// register repository
	genderRepo := repository.NewGenderRepo(db)
	provinceRepo := repository.NewProvinceRepo(db)
	cityRepo := repository.NewCityRepo(db)

	// register service
	genderService := service.NewGenderService(genderRepo)
	provinceService := service.NewProvinceService(provinceRepo, cityRepo)
	cityService := service.NewCityService(provinceRepo, cityRepo)

	// register handler
	genderHandler := handler.NewGenderHandler(genderService)
	provinceHandler := handler.NewProvinceHandler(provinceService)
	cityHandler := handler.NewCityHandler(cityService)

	r := gin.Default()
	r.NoRoute(gin.HandlerFunc(func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContextWithTracer(c, tracer, "handler no route")
		defer span.Finish()

		statusCode := http.StatusNotFound
		c.JSON(statusCode, &dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.GenerateStatusFromCode(statusCode),
			Message:    "endpoint not found",
		})
	}))

	v1 := r.Group("/api/v1")
	v1.GET("/test", func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c, "handler test")
		defer span.Finish()

		c.JSON(http.StatusOK, &dto.ApiResponse{
			StatusCode: http.StatusOK,
			Status:     "ok",
			Message:    "success test",
		})
		return
	})

	// gender routes
	router.CreateGenderRoutes(v1, genderHandler)

	// province routes
	router.CreateProvinceRoutes(v1, provinceHandler)

	// city routes
	router.CreateCityRoutes(v1, cityHandler)

	r.Run(":" + os.Getenv("APP_PORT"))
}
