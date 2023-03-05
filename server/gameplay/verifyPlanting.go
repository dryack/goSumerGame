package gameplay

import (
	"fmt"
)

func verifyPlanting(instructions Instructions, gameState GameState) error {
	const (
		ErrNegativeAcres       = "planting: you can't plant a negative amount of acres"
		ErrNotEnoughGrain      = "planting: you don't have enough grain to seed this many acres"
		ErrNotEnoughAcres      = "planting: you can't plant more acres than you own"
		ErrNotEnoughPopulation = "planting: you don't have enough population to seed this many acres"
	)

	ableToPlantByPopulation := gameState.Population * gameState.AcresPlantedToPopulationRatio
	ableToPlantByBushels := gameState.Bushels * gameState.AcresPlantedToBushelsRatio

	if instructions.PlantAcres < 0 {
		return fmt.Errorf("%s (%s: %d)", ErrNegativeAcres, InputString, instructions.PlantAcres)
	}
	if instructions.PlantAcres > gameState.Acres {
		return fmt.Errorf("%s (%s: %d, acres: %d)", ErrNotEnoughAcres, InputString, instructions.PlantAcres, gameState.Acres)
	}
	if instructions.PlantAcres > ableToPlantByBushels {
		return fmt.Errorf("%s (%s: %d, acres seeded per bushel of grain: %d)", ErrNotEnoughGrain, InputString, instructions.PlantAcres, gameState.AcresPlantedToBushelsRatio)
	}
	if instructions.PlantAcres > ableToPlantByPopulation {
		return fmt.Errorf("%s (%s: %d, acres planted per population: %d)", ErrNotEnoughPopulation, InputString, instructions.PlantAcres, gameState.AcresPlantedToPopulationRatio)
	}
	return nil
}
