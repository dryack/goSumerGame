package game

import (
	"encoding/gob"
	"errors"
	"fmt"
	"goSumerGame/server/model"
	"math/rand"
	"os"
	"strconv"
)

type GameState struct {
	Acres      int
	Population int
	Bushels    int
	BushelCost int
	DebugLvl   uint8
}

func (g *GameState) Initialize(debug uint8) error {
	g.DebugLvl = debug
	switch debug {
	case 0:
		g.BushelCost = 17 + rand.Intn(10)
	case 1:
		g.BushelCost = 20
	case 2:
		g.BushelCost = 20
	default:
		err := errors.New("invalid debug level:" + strconv.Itoa(int(debug)))
		return err
	}
	g.Bushels = 2000
	g.Population = 100
	g.Acres = 1000
	return nil
}

func (g *GameState) Save(game *model.Game) error {
	// TODO: Below commented code can probably be removed, but is kept as a reference for now
	// https://stackoverflow.com/questions/66966550/how-to-fetch-last-record-in-gorm
	//var gameDBID struct {
	//	ID int
	//}
	//database.Database.Table("games").Last(&gameDBID)

	// filename will be the user's id, an underscore, and then the game's expected id based on the games table last entry
	//filepath := "./saves/" + strconv.Itoa(int(game.UserID)) + "_" + strconv.Itoa(gameDBID.ID+1) + ".sav"

	fmt.Println(game.ID)
	filepath := "./saves/" + strconv.Itoa(int(game.UserID)) + "_" + strconv.Itoa(int(game.ID)) + ".sav"
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	game.Location = filepath
	encoder := gob.NewEncoder(file)
	encoder.Encode(g)
	return nil
}
