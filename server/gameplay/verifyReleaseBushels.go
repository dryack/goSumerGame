package gameplay

import (
	"fmt"
	"goSumerGame/server/model"
)

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
