package model

import (
	"goSumerGame/server/database"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Location string `gorm:"type:text" json:"location"`
	UserID   uint
	ID       uint
}

func (game *Game) Save() (*Game, error) {
	err := database.Database.Create(&game).Error
	if err != nil {
		return &Game{}, err
	}
	return game, nil
}

func (game *Game) Delete() (uint, error) {
	err := database.Database.Delete(&game).Error
	if err != nil {
		return game.ID, err
	}
	return game.ID, nil
}
