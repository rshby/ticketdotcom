package main

import (
	"backend/model/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getenv("database_name"))
}

func main() {
	fmt.Println("test")

	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &dto.ApiResponse{
			StatusCode: http.StatusOK,
			Status:     "ok",
			Message:    "success test",
		})
		return
	})

	r.Run(":" + os.Getenv("app_port"))
}
