package router

import (
	"backend/handler"
	"github.com/gin-gonic/gin"
)

func CreateProvinceRoutes(r *gin.RouterGroup, handler *handler.ProvinceHandler) *gin.RouterGroup {
	router := r.Group("")

	router.POST("/province", handler.Insert)

	return router
}
