package main

import (
	"quasar/api/routes"

	"github.com/gin-gonic/gin"
)

func Server() {
	r := gin.Default()

	routes.SetupPublicRoutes(r)

	r.Run()
}
