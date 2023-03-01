package gameplay

import (
	"errors"
	"math/rand"
	"strconv"
)

// GameState represents the state of the player's City State after a turn has been processed - or upon creation of the City State
type GameState struct {
	Year                 int
	Population           int
	PopulationFed        int
	PopulationFedByCows  int
	PopulationImmigrated int
	PopulationEmigrated  int
	PopulationStarved    int
	PopulationDied       int // needed?  we really may not need variables that can be derived by comparison
	Bushels              int
	AcreValue            int
	BushelYield          int
	PestsAte             int
	Acres                int
	AcresPlanted         int
	Granaries            int
	Plows                int
	Cows                 int
	Stelae               int
	BuildingPalace       int
	PalaceLvl1           bool
	PalaceLvl2           bool
	PalaceLvl3           bool
}

// Initialize accepts a debug level. It then sets an initial GameState of the player's City State, with the debug level determining some randomization
func (gs *GameState) Initialize(debug uint8) error {

	switch debug {
	case 0:
		gs.AcreValue = 17 + rand.Intn(10)
	case 1:
		gs.AcreValue = 20
	case 2:
		gs.AcreValue = 20
	default:
		err := errors.New("invalid debug level:" + strconv.Itoa(int(debug)))
		return err
	}
	gs.Year = 0
	gs.Population = 100
	gs.PopulationFed = 100
	gs.PopulationFedByCows = 0
	gs.PopulationImmigrated = 0
	gs.PopulationEmigrated = 0
	gs.PopulationStarved = 0
	gs.PopulationDied = 0
	gs.Bushels = 2000
	gs.BushelYield = 3
	gs.PestsAte = 200
	gs.Acres = 1000
	gs.AcresPlanted = 0
	gs.Granaries = 0
	gs.Plows = 0
	gs.Cows = 0
	gs.Stelae = 0
	gs.BuildingPalace = 0
	gs.PalaceLvl1 = false
	gs.PalaceLvl2 = false
	gs.PalaceLvl3 = false

	return nil
}
