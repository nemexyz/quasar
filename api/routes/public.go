package routes

import (
	"quasar/api/controller"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(gin *gin.Engine) {
	group := gin.Group("")

	group.POST("/topsecret", controller.PostTopSecret)
	group.POST("/topsecret_split/:satellite_name", controller.PostTopSecretSplit)
	group.GET("/topsecret_split", controller.GetTopSecretSplit)
}
