package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Location string `gorm:"type:text" json:"location"`
	UserID   uint
}
