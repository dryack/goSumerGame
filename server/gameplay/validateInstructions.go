package gameplay

import (
	"errors"
	"fmt"
	"goSumerGame/server/model"
)

// validateInstructions ensures that all player instructions are valid, and then
// simulates them being applied in order, to ensure that the entire sequence of
// orders will not result in an invalid state. Any error will cause the turn to
// not be taken, and the specific error will be returned to the player by gin.
// newGameState is the candidate gameState to be appended to the
// gameSession.History, assuming there are no errors.
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

	return nil
}

func verifyPurchaseAcres(instructions model.Instructions, gameState GameState) error {
	totalAcres := gameState.Acres + instructions.PurchaseAcres
	if totalAcres < 0 {
		err := errors.New("you don't have enough acres to sell")
		err = fmt.Errorf("%s (acres: %d, change: %d)", err.Error(), gameState.Acres, instructions.PurchaseAcres)
		return err
	}
	totalCost := instructions.PurchaseAcres * gameState.AcreValue
	if totalCost > gameState.Bushels {
		err := errors.New("you don't have enough bushels to afford this")
		err = fmt.Errorf("%s (bushels: %d, change: %d, total cost: %d)", err.Error(), gameState.Bushels, instructions.PurchaseAcres, totalCost)
		return err
	}
	return nil
}

func verifyReleaseBushels(instructions model.Instructions, gameState GameState) error {
	if instructions.ReleaseBushels < 0 {
		err := errors.New("you can't disburse a negative amount")
		err = fmt.Errorf("%s (change: %d)", err.Error(), instructions.ReleaseBushels)
		return err
	}
	if instructions.ReleaseBushels > gameState.Bushels {
		err := errors.New("you don't have that many bushels")
		err = fmt.Errorf("%s (bushels: %d, change: %d)", err.Error(), gameState.Bushels, instructions.ReleaseBushels)
		return err
	}
	return nil
}
