package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pyrtos.com/api/models"
)

// GET /devices
func FindDevices(c *gin.Context) {
	var devices []models.Device
	models.DB.Find(&devices)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// POST /devices
// Create new device
func CreateDevice(c *gin.Context) {
	// Validate input
	var input models.CreateDeviceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create device
	device := models.Device{MAC: input.MAC, Name: input.Name}
	models.DB.Create(&device)

	c.JSON(http.StatusOK, gin.H{"data": device})
}
