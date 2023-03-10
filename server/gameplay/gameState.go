package gameplay

import (
	"errors"
	"math/rand"
	"strconv"
)

// GameState represents the state of the player's City State after a turn has been processed - or upon creation of the City State
type GameState struct {
	Year                          int  `json:"year"`
	Population                    int  `json:"population"`
	PopulationBorn                int  `json:"population_born"`
	PopulationFed                 int  `json:"population_fed"`
	PopulationFedByCows           int  `json:"population_fed_by_cows"`
	PopulationImmigrated          int  `json:"population_immigrated"`
	PopulationEmigrated           int  `json:"population_emigrated"`
	PopulationStarved             int  `json:"population_starved"`
	PopulationDied                int  `json:"population_died"` // needed?  we really may not need variables that can be derived by comparison
	Bushels                       int  `json:"bushels"`
	BushelsReleased               int  `json:"bushels_released"`
	BushelsToFeedAPop             int  `json:"bushels_to_feed_a_pop"`
	AcreValue                     int  `json:"acre_value"`
	AcreGrainYield                int  `json:"acre_grain_yield"`
	PestsAte                      int  `json:"pests_ate"`
	Acres                         int  `json:"acres"`
	AcresPlanted                  int  `json:"acres_planted"`
	AcresPlantedToBushelsRatio    int  `json:"acres_planted_to_bushels_ratio"`
	AcresPlantedToPopulationRatio int  `json:"acres_planted_to_population_ratio"`
	Granaries                     int  `json:"granaries"`
	Plows                         int  `json:"plows"`
	CowHerds                      int  `json:"cow_herds"`
	Stelae                        int  `json:"stelae"`
	Ziggurats                     int  `json:"ziggurats"`
	Temples                       int  `json:"temples"`
	Plague                        int  `json:"plague"`
	BuildingPalace                int  `json:"building_palace"`
	PalaceLvl1                    bool `json:"palace_lvl_1"`
	PalaceLvl2                    bool `json:"palace_lvl_2"`
	PalaceLvl3                    bool `json:"palace_lvl_3"`
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
	gs.PopulationBorn = 0
	gs.PopulationFed = 100
	gs.PopulationFedByCows = 0
	gs.PopulationImmigrated = 0
	gs.PopulationEmigrated = 0
	gs.PopulationStarved = 0
	gs.PopulationDied = 0
	gs.Bushels = StartingBushels
	gs.BushelsReleased = StartingBushelsReleased
	gs.BushelsToFeedAPop = StartingBushelsToFeedAPop
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
