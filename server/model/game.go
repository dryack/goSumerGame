package model

import (
	"goSumerGame/server/database"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Location string `gorm:"type:text" json:"location"`
	UserID   uint
}

func (game *Game) Save() (*Game, error) {
	err := database.Database.Create(&game).Error
	if err != nil {
		return &Game{}, err
	}
	return game, nil
}

// game.Delete() sets the specified game ID to deleted and returns the number of rows impacted. If this number is 0 without an error, it implies the Game.ID was not found.
func (game *Game) Delete() (int64, error) {
	result := database.Database.Delete(&game)
	err := result.Error
	if err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
