package gameplay

import (
	"fmt"
	"goSumerGame/server/model"
)

const (
	InputString = "input" // saving some typing across a lot of error messages
)

// validateInstructions ensures that all player instructions are valid, and then
// simulates them being applied in order, to ensure that the entire sequence of
// orders will not result in an invalid state.
//
// - Any error will cause the turn to not be taken, and the specific error will be
// returned to the player by gin. newGameState is the candidate gameState to be
// appended to the gameSession.History, assuming there are no errors.
//
// - tempGameState is the prior turn, used to perform operations in order,
// for simulating the player's instructions
func validateInstructions(instructions model.Instructions, gameState GameState, newGameState *GameState) error {
	tempGameState := gameState

	// TODO: Blocks after a verifySomething() should be extracted to their own functions
	err := verifyPurchaseAcres(instructions, tempGameState)
	if err != nil {
		return err
	}
	newGameState.Acres += instructions.PurchaseAcres
	newGameState.Bushels -= instructions.PurchaseAcres * gameState.AcreValue
	tempGameState.Acres = newGameState.Acres
	tempGameState.Bushels = newGameState.Bushels

	err = verifyReleaseBushels(instructions, tempGameState)
	if err != nil {
		return err
	}
	newGameState.Bushels -= instructions.ReleaseBushels
	tempGameState.Bushels = newGameState.Bushels

	err = verifyPlanting(instructions, tempGameState)
	if err != nil {
		return err
	}
	newGameState.Bushels -= stochasticRoundDivide(instructions.PlantAcres, gameState.AcresPlantedToBushelsRatio)
	newGameState.AcresPlanted = instructions.PlantAcres
	tempGameState.Bushels = newGameState.Bushels

	return nil
}

func verifyPurchaseAcres(instructions model.Instructions, gameState GameState) error {
	const (
		ErrNotEnoughAcres   = "purchase/sell acres: you don't have enough acres to sell"
		ErrNotEnoughBushels = "purchase/sell acres: you don't have enough bushels to afford this"
	)
	totalAcres := gameState.Acres + instructions.PurchaseAcres
	if totalAcres < 0 {
		return fmt.Errorf("%s (%s: %d, change: %d)", ErrNotEnoughAcres, InputString, gameState.Acres, instructions.PurchaseAcres)
	}
	totalCost := instructions.PurchaseAcres * gameState.AcreValue
	if totalCost > gameState.Bushels {
		return fmt.Errorf("%s (%s: %d, change: %d, total cost: %d)", ErrNotEnoughBushels, InputString, gameState.Bushels, instructions.PurchaseAcres, totalCost)
	}
	return nil
}

func verifyReleaseBushels(instructions model.Instructions, gameState GameState) error {
	const (
		ErrNegativeBushels  = "releasing food: you can't disburse a negative amount of bushels"
		ErrNotEnoughBushels = "releasing food: you don't have that many bushels"
	)
	if instructions.ReleaseBushels < 0 {
		return fmt.Errorf("%s (%s: %d)", ErrNegativeBushels, InputString, instructions.ReleaseBushels)
	}
	if instructions.ReleaseBushels > gameState.Bushels {
		return fmt.Errorf("%s (%s: %d, change: %d)", ErrNotEnoughBushels, InputString, gameState.Bushels, instructions.ReleaseBushels)
	}
	return nil
}

func verifyPlanting(instructions model.Instructions, gameState GameState) error {
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
