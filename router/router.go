package router

import (
	"web/controller"
	"github.com/gin-gonic/gin"
)

func Router(){
	router := gin.Default()
	router.POST("/user", controller.UserSaveHandler)
	router.PUT("/user/:id", controller.UserUpdateHandler)
	router.DELETE("/user/:id", controller.UserDeleteHandler)
	router.Run()
}