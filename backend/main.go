package main

import (
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

	fmt.Println(os.Getenv("DATABASE_NAME"))
}

func main() {
	fmt.Println("test")

	// register tracing
	tracer, _, err := tracing.ConnectJaegerTracing()
	//defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		panic("err cant connect jaeger : " + err.Error())
	}

	r := gin.Default()
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

	r.Run(":" + os.Getenv("APP_PORT"))
}
