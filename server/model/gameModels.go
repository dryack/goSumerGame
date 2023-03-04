package model

import (
	"goSumerGame/server/database"
	"gorm.io/gorm"
)

// Game represents an individual row in the games table
type Game struct {
	gorm.Model
	UserID uint
	// Debug game levels:
	// 0 = normal game
	// 1 = starting conditions always the same
	// 2 = all random conditions are the same
	// etc.
	Debug    uint8 `gorm:"type:smallint" json:"debug"`
	Turns    uint
	Location string
	Note     string `gorm:"type:varchar(600)" json:"note"`
}

// Instructions represent the player's instructions for a game turn
type Instructions struct {
	GameID            uint
	PurchaseAcres     int
	ReleaseBushels    int
	PlantAcres        int
	PurchaseGranaries int
}

// Save saves a Game to the games table
func (game *Game) Save() (*Game, error) {
	err := database.Database.Create(&game).Error
	if err != nil {
		return &Game{}, err
	}
	return game, nil
}

// Update upserts a Game to the games table
func (game *Game) Update() (*Game, error) {
	err := database.Database.Updates(&game).Error
	if err != nil {
		return &Game{}, err
	}
	return game, nil
}

// Delete sets the specified game ID to deleted and returns the number of rows impacted. If this number is 0 without an error, it implies the Game.ID was not found.
func (game *Game) Delete() (int64, error) {
	result := database.Database.Delete(&game)
	err := result.Error
	if err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
