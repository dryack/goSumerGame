package gameplay

import (
	"encoding/gob"
	"errors"
	"fmt"
	"goSumerGame/server/model"
	"os"
	"strconv"
)

// GameSession is the base data structure represent a single ongoing play session
type GameSession struct {
	Meta    GameMeta
	History GameHistory
}

// GameHistory contains every GameState produced during a given play session
type GameHistory []*GameState

// GameMeta contains GameSession information that isn't related to the GameState
type GameMeta struct {
	DebugLevel uint8
}

// Initialize accepts a debug level and a pointer to an empty GameState, it pushes the GameState onto its History slice as the initial GameState of the play session
func (g *GameSession) Initialize(debug uint8, state *GameState) error {
	g.Meta.DebugLevel = debug
	err := state.Initialize(debug)
	if err != nil {
		return err
	}
	g.History = append(g.History, state)
	return nil
}

// Save accepts a pointer to a *model.Game, whose fields are used to determine the name of the .sav file when a GameSession is saved to disk
func (g *GameSession) Save(game *model.Game) error {
	// TODO: Below commented code can probably be removed, but is kept as a reference for now
	// https://stackoverflow.com/questions/66966550/how-to-fetch-last-record-in-gorm
	// var gameDBID struct {
	//	ID int
	// }
	// database.Database.Table("games").Last(&gameDBID)

	// filename will be the user's id, an underscore, and then the game's expected id based on the games table last entry
	// filepath := "./saves/" + strconv.Itoa(int(game.UserID)) + "_" + strconv.Itoa(gameDBID.ID+1) + ".sav"

	saveDir := "./saves/"
	saveExtension := ".sav"
	saveUserID := strconv.Itoa(int(game.UserID))
	saveGameDBID := strconv.Itoa(int(game.ID))
	// check if the saveDir exists - create it if not
	if _, err := os.Stat(saveDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(saveDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
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

func (g *GameSession) Delete(gameLocation string) error {
	fmt.Println(gameLocation)
	err := os.Remove(gameLocation)
	if err != nil {
		return err
	}
	return nil
}

// Load accepts a pointer to a model.Game, from which it uses the Location field to determine where a GameSession is stored on disk, after which loads and decodes that .sav file
func (g *GameSession) Load(gameLocation string) error {
	file, err := os.Open(gameLocation)
	defer file.Close()
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(g)
	if err != nil {
		return err
	}
	return nil
}

func (g *GameSession) Test(instructions *Instructions, game *model.Game) (TurnResponse, error) {
	oldGameState := g.History[len(g.History)-1]
	newGameState := oldGameState
	err := validateInstructions(*instructions, *oldGameState, newGameState)
	if err != nil {
		return TurnResponse{}, err
	}

	g.History = append(g.History, newGameState)
	err = g.Save(game)
	if err != nil {
		return TurnResponse{}, err
	}
	return TurnResponse{
		GameState: *newGameState,
		Messages:  nil,
	}, nil
}

func (g *GameSession) RunTurn(instructions *Instructions, game *model.Game) (TurnResponse, error) {
	oldGameState := g.History[len(g.History)-1]
	newGameState := oldGameState
	var messages []string
	err := validateInstructions(*instructions, *oldGameState, newGameState)
	if err != nil {
		return TurnResponse{}, err
	}

	err = processTurn(*oldGameState, newGameState, &messages)
	if err != nil {
		return TurnResponse{}, err
	}

	g.History = append(g.History, newGameState)
	err = g.Save(game)
	if err != nil {
		return TurnResponse{}, err
	}
	return TurnResponse{
		GameState: *newGameState,
		Messages:  messages,
	}, nil
}
