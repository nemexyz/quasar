package routes

import (
	"quasar/api/controller"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(gin *gin.Engine) {
	group := gin.Group("")

	group.POST("/topsecret", controller.PostTopSecret)
}
