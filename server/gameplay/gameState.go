package gameplay

import (
	"errors"
	"math/rand"
	"strconv"
)

type GameState struct {
	Year       int
	Population int
	Bushels    int
	BushelCost int
	Acres      int
}

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
