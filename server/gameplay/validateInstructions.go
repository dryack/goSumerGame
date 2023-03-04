package gameplay

import (
	"goSumerGame/server/model"
)

const (
	InputString = "input" // saving some typing across a lot of error messages
)

// validateInstructions ensures that all player instructions are valid. It
// simulates them being applied in order, to ensure that the entire sequence of
// orders will not result in an invalid state.
//
// - Any error will cause the turn to not be taken, and the specific error will be
// returned to the player by gin.
//
// - tempGameState is the prior turn's gameState. Instructions from the player
// are checked against and applied to items like Acres and Bushels, ensuring
// commands are not invalid.
//
// - newGameState is a pointer to the candidate gameState that will be appended
// to the gameSession.History if there are no errors.
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

	err = verifyGranaries(instructions, tempGameState)
	if err != nil {
		return err
	}
	newGameState.Bushels -= instructions.PurchaseGranaries * CostGranary
	newGameState.Granaries += instructions.PurchaseGranaries

	return nil
}
