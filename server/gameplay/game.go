package gameplay

import (
	"encoding/gob"
	"goSumerGame/server/model"
	"os"
	"strconv"
)

type Game struct {
	Meta    GameMeta
	History GameHistory
}

type GameHistory []*GameState

type GameMeta struct {
	DebugLevel uint8
}

func (g *Game) Initialize(debug uint8, state *GameState) error {
	g.Meta.DebugLevel = debug
	err := state.Initialize(debug)
	if err != nil {
		return err
	}
	g.History = append(g.History, state)
	return nil
}

func (g *Game) Save(game *model.Game) error {
	// TODO: Below commented code can probably be removed, but is kept as a reference for now
	// https://stackoverflow.com/questions/66966550/how-to-fetch-last-record-in-gorm
	//var gameDBID struct {
	//	ID int
	//}
	//database.Database.Table("games").Last(&gameDBID)

	// filename will be the user's id, an underscore, and then the game's expected id based on the games table last entry
	//filepath := "./saves/" + strconv.Itoa(int(game.UserID)) + "_" + strconv.Itoa(gameDBID.ID+1) + ".sav"
	saveDir := "./saves/"
	saveExtension := ".sav"
	saveUserID := strconv.Itoa(int(game.UserID))
	saveGameDBID := strconv.Itoa(int(game.ID))
	filepath := saveDir + saveUserID + "_" + saveGameDBID + saveExtension
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

func (g *Game) Load(game *model.Game) error {
	filepath := game.Location

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(g)
	return nil
}
