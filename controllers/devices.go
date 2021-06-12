package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pyrtos.com/api/models"
)

type CreateDeviceInput struct {
	MAC_address string `json:"mac_address" binding:"required"`
	Name        string `json:"name"`
}

type UpdateDeviceInput struct {
	MAC_address string `json:"mac_address"`
	Name        string `json:"name"`
}

// GET /devices
func FindDevices(c *gin.Context) {
	var devices []models.Device
	models.DB.Find(&devices)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

// GET /devices/:id
func FindDevice(c *gin.Context) { // Get model if exist
	var device models.Device

	if err := models.DB.Where("id = ?", c.Param("id")).First(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": device})
}

// POST /devices
// Create new device
func CreateDevice(c *gin.Context) {
	// Validate input
	var input CreateDeviceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create device
	device := models.Device{MAC_address: input.MAC_address, Name: input.Name}
	models.DB.Create(&device)

	c.JSON(http.StatusOK, gin.H{"data": device})
}

// PATCH /devices/:id
// Update a device
func UpdateDevice(c *gin.Context) {
	var device models.Device
	if err := models.DB.Where("id = ?", c.Param("id")).First(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateDeviceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&device).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": device})
}

// DELETE /devices/:id
// Delete a device
func DeleteDevice(c *gin.Context) {
	var device models.Device
	if err := models.DB.Where("id = ?", c.Param("id")).First(&device).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&device)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
