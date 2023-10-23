package router

import (
	"backend/handler"
	"github.com/gin-gonic/gin"
)

func CreateCityRoutes(r *gin.RouterGroup, handler *handler.CityHandler) *gin.RouterGroup {
	route := r.Group("")

	route.GET("/cities", handler.GetAllCity)
	route.POST("/city", handler.Insert)
	route.GET("/province/cities/:provinceId", handler.GetByProvinceId)
	return route
}
