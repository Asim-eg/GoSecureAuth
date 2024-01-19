package routes

import (
	"github.com/GoSecureAuth/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.Group("/GoAuthenticate")
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
}
