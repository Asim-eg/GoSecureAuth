package routes

import (
	"github.com/GoSecureAuth/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Group("/GoAuthenticate")
	router.GET("/users", controller.GetUsers())
	router.GET("/users/:id", controller.GetUser())
}
