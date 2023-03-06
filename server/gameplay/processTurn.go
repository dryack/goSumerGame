package gameplay

import "math/rand"

func processTurn(oldGameState GameState, newGameState *GameState, messages *[]string) error {
	newGameState.Year += 1
	msg := doPlague(newGameState)
	if msg != "" {
		*messages = append(*messages, msg)
	}
	newGameState.PopulationFed = stochasticRoundDivide(newGameState.BushelsReleased, StartingBushelsToFeedAPop)
	newGameState.AcreGrainYield = rand.Intn(9) + 1
	newGameState.AcreValue = 17 + rand.Intn(10)

	return nil
}

func doPlague(gameState *GameState) string {
	// TODO: Make plagues more likely based on sanitation, population, and tech
	const (
		NewPlague     = "A plague has struck!"
		OngoingPlague = "The plague continues!"
	)

	// no plague
	if gameState.Year <= 0 {
		return ""
	}
	if rand.Intn(19)-gameState.Plague > 0 {
		gameState.Plague = 0
		return ""
	}

	// Plague begets more plague
	gameState.Plague += 1
	// Avoid permanent plagues
	if gameState.Plague > 10 {
		gameState.Plague = 10
	}
	if gameState.Plague > 1 {
		return OngoingPlague
	}
	return NewPlague
}
