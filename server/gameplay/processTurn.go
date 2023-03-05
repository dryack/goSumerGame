package gameplay

import "math/rand"

func processTurn(oldGameState GameState, newGameState *GameState) error {
	newGameState.Year += 1
	doPlague(newGameState)
	return nil
}

func doPlague(gameState *GameState) {
	// TODO: Make plagues more likely based on sanitation, population, and tech

	// no plague
	if gameState.Year <= 0 {
		return
	}
	if rand.Intn(19)-gameState.Plague > 0 {
		gameState.Plague = 0
		return
	}

	// Plague begets more plague
	gameState.Plague += 1
	// Avoid permanent plagues
	if gameState.Plague > 10 {
		gameState.Plague = 10
	}
}
