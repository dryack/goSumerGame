package gameplay

import (
	"errors"
	"fmt"
	"goSumerGame/server/model"
	"sort"
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
	totalAcres := gameState.Acres + instructions.PurchaseAcres
	if totalAcres < 0 {
		err := errors.New("purchase/sell acres: you don't have enough acres to sell")
		err = fmt.Errorf("%s (acres: %d, change: %d)", err.Error(), gameState.Acres, instructions.PurchaseAcres)
		return err
	}
	totalCost := instructions.PurchaseAcres * gameState.AcreValue
	if totalCost > gameState.Bushels {
		err := errors.New("purchase/sell acres: you don't have enough bushels to afford this")
		err = fmt.Errorf("%s (bushels: %d, change: %d, total cost: %d)", err.Error(), gameState.Bushels, instructions.PurchaseAcres, totalCost)
		return err
	}
	return nil
}

func verifyReleaseBushels(instructions model.Instructions, gameState GameState) error {
	if instructions.ReleaseBushels < 0 {
		err := errors.New("releasing food: you can't disburse a negative amount of bushels")
		err = fmt.Errorf("%s (change: %d)", err.Error(), instructions.ReleaseBushels)
		return err
	}
	if instructions.ReleaseBushels > gameState.Bushels {
		err := errors.New("releasing food: you don't have that many bushels")
		err = fmt.Errorf("%s (bushels: %d, change: %d)", err.Error(), gameState.Bushels, instructions.ReleaseBushels)
		return err
	}
	return nil
}

func verifyPlanting(instructions model.Instructions, gameState GameState) error {
	ableToPlantByPopulation := gameState.Population * gameState.AcresPlantedToPopulationRatio
	ableToPlantByBushels := gameState.Bushels * gameState.AcresPlantedToBushelsRatio
	var limitingFactorForPlanting = []int{ableToPlantByBushels, ableToPlantByPopulation, gameState.Acres}
	sort.Ints(limitingFactorForPlanting)

	var errorMsg string
	switch limitingFactorForPlanting[0] {
	case gameState.Bushels:
		errorMsg = "planting: you don't have enough grain to seed this many acres"
	case ableToPlantByPopulation:
		errorMsg = "planting: you can't plant more acres than you own"
	case ableToPlantByBushels:
		errorMsg = "planting: you don't have enough grain to seed this many acres"
	}

	if instructions.PlantAcres < 0 {
		err := errors.New("planting: you can't plant a negative amount of acres")
		err = fmt.Errorf("%s (planting: %d)", err.Error(), instructions.PlantAcres)
		return err
	}
	if instructions.PlantAcres > gameState.Acres {
		err := errors.New(errorMsg)
		err = fmt.Errorf("%s (planting: %d, acres: %d)", err.Error(), instructions.PlantAcres, gameState.Acres)
		return err
	}
	if instructions.PlantAcres/gameState.AcresPlantedToBushelsRatio > gameState.Bushels {
		err := errors.New(errorMsg)
		err = fmt.Errorf("%s (planting: %d, acres seeded per bushel of grain: %d)", err.Error(), instructions.PlantAcres, gameState.AcresPlantedToBushelsRatio)
		return err
	}
	if instructions.PlantAcres/gameState.AcresPlantedToPopulationRatio < gameState.Population {
		err := errors.New(errorMsg)
		err = fmt.Errorf("%s (planting: %d, acres planted per population: %d)", err.Error(), instructions.PlantAcres, gameState.AcresPlantedToPopulationRatio)
	}
	return nil
}
