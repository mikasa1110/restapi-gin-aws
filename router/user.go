package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mikasa1110/restapi-gin-gorm-docker/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}
