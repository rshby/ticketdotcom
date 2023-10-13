package main

import (
	"backend/database"
	"backend/model/dto"
	"backend/tracing"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("app run")

	db := database.ConnectDB()
	fmt.Println(db)

	// register tracing
	tracer, _, err := tracing.ConnectJaegerTracing()
	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		panic("err cant connect jaeger : " + err.Error())
	}

	r := gin.Default()
	r.NoRoute(gin.HandlerFunc(func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContextWithTracer(c, tracer, "handler no route")
		defer span.Finish()

		c.JSON(http.StatusNotFound, &dto.ApiResponse{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Message:    "endpoint not found",
		})
	}))

	v1 := r.Group("/api/v1")
	v1.GET("/test", func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c, "handler test")
		defer span.Finish()

		// call function from service
		// serviec.GetAllData(ctx)

		//

		c.JSON(http.StatusOK, &dto.ApiResponse{
			StatusCode: http.StatusOK,
			Status:     "ok",
			Message:    "success test",
		})
		return
	})

	r.Run(":" + os.Getenv("APP_PORT"))
}
