package gameplay

import (
	"fmt"
	"goSumerGame/server/model"
)

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
