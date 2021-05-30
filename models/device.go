package models

type Device struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	MAC  string `json:"mac_address"`
	Name string `json:"name"`
}

type CreateDeviceInput struct {
	MAC  string `json:"mac_address" binding:"required"`
	Name string `json:"name"`
}
