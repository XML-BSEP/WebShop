package domain

import "gorm.io/gorm"

type AdminToken struct {
	gorm.Model
	Token string `json:"token"`
	UserId string `json:"userId"`
}
