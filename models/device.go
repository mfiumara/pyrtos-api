package models

type Device struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	MAC_address string `json:"mac_address"`
	Name        string `json:"name"`
}
