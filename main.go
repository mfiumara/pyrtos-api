package main

import (
	"github.com/gin-gonic/gin"
	"pyrtos.com/api/controllers"
	"pyrtos.com/api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/devices", controllers.FindDevices)
	r.POST("/devices", controllers.CreateDevice)

	r.Run()
}
