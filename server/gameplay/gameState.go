package gameplay

import (
	"errors"
	"math/rand"
	"strconv"
)

// GameState represents the state of the player's City State after a turn has been processed - or upon creation of the City State
type GameState struct {
	Year       int
	Population int
	Bushels    int
	BushelCost int
	Acres      int
}

// Initialize accepts a debug level. It then sets an initial GameState of the player's City State, with the debug level determining some randomization
func (gs *GameState) Initialize(debug uint8) error {

	switch debug {
	case 0:
		gs.BushelCost = 17 + rand.Intn(10)
	case 1:
		gs.BushelCost = 20
	case 2:
		gs.BushelCost = 20
	default:
		err := errors.New("invalid debug level:" + strconv.Itoa(int(debug)))
		return err
	}
	gs.Bushels = 2000
	gs.Population = 100
	gs.Acres = 1000
	return nil
}
