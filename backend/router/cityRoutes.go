package router

import (
	"backend/handler"
	"github.com/gin-gonic/gin"
)

func CreateCityRoutes(r *gin.RouterGroup, handler *handler.CityHandler) *gin.RouterGroup {
	route := r.Group("")

	route.GET("/cities", handler.GetAllCity)
	return route
}
