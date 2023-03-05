package gameplay

import (
	"errors"
	"math/rand"
	"strconv"
)

// GameState represents the state of the player's City State after a turn has been processed - or upon creation of the City State
type GameState struct {
	Year                          int  `json:"year,omitempty"`
	Population                    int  `json:"population,omitempty"`
	PopulationFed                 int  `json:"population_fed,omitempty"`
	PopulationFedByCows           int  `json:"population_fed_by_cows,omitempty"`
	PopulationImmigrated          int  `json:"population_immigrated,omitempty"`
	PopulationEmigrated           int  `json:"population_emigrated,omitempty"`
	PopulationStarved             int  `json:"population_starved,omitempty"`
	PopulationDied                int  `json:"population_died,omitempty"` // needed?  we really may not need variables that can be derived by comparison
	Bushels                       int  `json:"bushels,omitempty"`
	AcreValue                     int  `json:"acre_value,omitempty"`
	AcreGrainYield                int  `json:"acre_grain_yield,omitempty"`
	PestsAte                      int  `json:"pests_ate,omitempty"`
	Acres                         int  `json:"acres,omitempty"`
	AcresPlanted                  int  `json:"acres_planted,omitempty"`
	AcresPlantedToBushelsRatio    int  `json:"acres_planted_to_bushels_ratio,omitempty"`
	AcresPlantedToPopulationRatio int  `json:"acres_planted_to_population_ratio,omitempty"`
	Granaries                     int  `json:"granaries,omitempty"`
	Plows                         int  `json:"plows,omitempty"`
	CowHerds                      int  `json:"cow_herds,omitempty"`
	Stelae                        int  `json:"stelae,omitempty"`
	Ziggurats                     int  `json:"ziggurats,omitempty"`
	Temples                       int  `json:"temples,omitempty"`
	Plague                        int  `json:"plague,omitempty"`
	BuildingPalace                int  `json:"building_palace,omitempty"`
	PalaceLvl1                    bool `json:"palace_lvl_1,omitempty"`
	PalaceLvl2                    bool `json:"palace_lvl_2,omitempty"`
	PalaceLvl3                    bool `json:"palace_lvl_3,omitempty"`
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
	gs.Population = StartingPopulation
	gs.PopulationFed = 100
	gs.PopulationFedByCows = 0
	gs.PopulationImmigrated = 0
	gs.PopulationEmigrated = 0
	gs.PopulationStarved = 0
	gs.PopulationDied = 0
	gs.Bushels = StartingBushels
	gs.AcreGrainYield = StartingAcreGrainYield
	gs.PestsAte = 200
	gs.Acres = StartingAcres
	gs.AcresPlanted = 0
	gs.AcresPlantedToBushelsRatio = StartingBushelSeedingRatio // every bushel can seed 2 acres with starting tech
	gs.AcresPlantedToPopulationRatio = StartingPopulationRatio // every pop can plant 10 acres with starting tech
	gs.Granaries = 0
	gs.Plows = 0
	gs.CowHerds = 0
	gs.Stelae = 0
	gs.Ziggurats = 0
	gs.Temples = 0
	gs.Plague = 0
	gs.BuildingPalace = 0
	gs.PalaceLvl1 = false
	gs.PalaceLvl2 = false
	gs.PalaceLvl3 = false

	return nil
}
