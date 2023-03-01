package gameplay

import (
	"errors"
	"fmt"
	"goSumerGame/server/model"
)

func validateInstructions(instructions model.Instructions, gameState GameState) error {
	err := verifyPurchaseAcres(instructions, gameState)
	if err != nil {
		return err
	}
	return nil
}

func verifyPurchaseAcres(instructions model.Instructions, gameState GameState) error {
	totAcres := gameState.Acres + instructions.PurchaseAcres
	if totAcres < 0 {
		err1 := errors.New("you don't have enough acres to sell")
		err2 := fmt.Errorf(" acres: %d", gameState.Acres)
		err3 := fmt.Errorf(" change: %d", instructions.PurchaseAcres)
		return errors.Join(err1, err2, err3)
	}
	totalCost := instructions.PurchaseAcres * gameState.AcreValue
	if totalCost > gameState.Bushels {
		err1 := errors.New("you don't have enough bushels to afford this")
		err2 := fmt.Errorf(" bushels: %d", gameState.Bushels)
		err3 := fmt.Errorf(" change: %d", instructions.PurchaseAcres)
		err4 := fmt.Errorf(" total cost: %d", totalCost)
		return errors.Join(err1, err2, err3, err4)
	}
	return nil
}
