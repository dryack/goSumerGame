package gameplay

import (
	"fmt"
	"goSumerGame/server/model"
)

func verifyGranaries(instructions model.Instructions, gameState GameState) error {
	const (
		ErrNegativeGranaries = "granaries: you can't purchase a negative number of granaries"
		ErrNotEnoughBushels  = "granaries: you don't have enough bushels to purchase these granaries"
	)
	maxGranaries := gameState.Bushels / CostGranary

	if instructions.PurchaseGranaries < 0 {
		return fmt.Errorf("%s (%s: %d)", ErrNegativeGranaries, InputString, instructions.PurchaseGranaries)
	}
	if instructions.PurchaseGranaries > maxGranaries {
		return fmt.Errorf("%s (%s: %d, bushels: %d", ErrNotEnoughBushels, InputString, instructions.PurchaseGranaries, gameState.Bushels)
	}
	return nil
}
