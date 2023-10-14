package router

import (
	"backend/handler"
	"github.com/gin-gonic/gin"
)

func CreateGenderRoutes(r *gin.RouterGroup, handler *handler.GenderHandler) *gin.RouterGroup {
	router := r.Group("")

	router.POST("/gender", handler.Insert)
	router.GET("/genders", handler.GetAll)

	return router
}
