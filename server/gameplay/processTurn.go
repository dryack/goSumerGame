package gameplay

import "math/rand"

func processTurn(oldGameState GameState, newGameState *GameState, messages *[]string) error {
	newGameState.Year += 1
	randomGenerator := realRandomizer{}
	msg := doPlague(newGameState, randomGenerator)
	if msg != "" {
		*messages = append(*messages, msg)
	}

	r := realRandomizer{}
	newGameState.Population, msg = doPopulation(oldGameState, newGameState, messages, r)
	if msg != "" {
		*messages = append(*messages, msg)
	}

	newGameState.AcreGrainYield = rand.Intn(9) + 1
	newGameState.AcreValue = 17 + rand.Intn(10)

	return nil
}

func doPopulation(oldGameState GameState, newGameState *GameState, messages *[]string, r randomizer) (int, string) {
	const (
		Starvation = "There is starvation in the city!"
	)
	var msg string

	newGameState.PopulationFed = stochasticRoundDivide(newGameState.BushelsReleased, StartingBushelsToFeedAPop)
	newGameState.PopulationBorn = stochasticRoundDivide(oldGameState.Population, r.Intn(8)+2)
	newGameState.PopulationStarved = zeroIfNegative(oldGameState.Population - newGameState.PopulationFed)
	if newGameState.PopulationStarved > 0 {
		msg = Starvation
	}
	// TODO:  immigration and emigration should be impacted by general moral, overfeeding, and possibly other things
	// immigration and emigration numbers established here will be overridden below in cases of Plague
	newGameState.PopulationImmigrated = int(0.1 * float64(r.Intn(oldGameState.Population)+1))
	newGameState.PopulationEmigrated = int(0.075 * float64(r.Intn(oldGameState.Population)+1))

	if newGameState.Plague > 0 {
		newGameState.PopulationBorn = zeroIfNegative(newGameState.PopulationBorn/2) - (newGameState.Plague * 2)

		newGameState.PopulationImmigrated = 0
		newGameState.PopulationEmigrated = int(0.1 * float64(r.Intn(oldGameState.Population)+newGameState.Plague))
	}
	finalPopulation := (oldGameState.Population + newGameState.PopulationBorn + newGameState.PopulationImmigrated) - (newGameState.PopulationStarved + newGameState.PopulationEmigrated)
	return finalPopulation, msg
}

func doPlague(gameState *GameState, r randomizer) string {
	// TODO: Make plagues more likely based on sanitation, population, and tech
	const (
		NewPlague     = "A plague has struck!"
		OngoingPlague = "The plague continues!"
		EndOfPlague   = "The plague has ended!"
	)

	// no plague
	if gameState.Year <= 0 {
		return ""
	}
	if r.Intn(19)-gameState.Plague > 0 {
		if gameState.Plague > 0 {
			gameState.Plague = 0
			return EndOfPlague
		}
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

func zeroIfNegative(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

// This stuff courtesy of an extended conversation with ChatGPT, that proved neither of us have any idea how to do
// this sort of stuff without help. The point is to allow completely deterministic "random" numbers when writing tests.
// sequentialFixedRandomizer was my own idea, as I discovered that fixedRandomizer wasn't enough to meet my needs.

type randomizer interface {
	Intn(n int) int
}

// A realRandomizer uses rand.Intn to return a random number between and N
type realRandomizer struct{}

// A fixedRandomizer returns a specified integer each time it is called. It is used to permit unit testing of functions
// that rely on a rand.Intn
type fixedRandomizer struct {
	value int
}

// A sequentialFixedRandomizer returns an integer from values each time it is called. It is used to permit unit testing of functions
// that rely on a rand.Intn
type sequentialFixedRandomizer struct {
	called uint
	values []int
}

func (r realRandomizer) Intn(n int) int {
	return rand.Intn(n)
}

func (r fixedRandomizer) Intn(n int) int {
	return r.value
}

func (r *sequentialFixedRandomizer) Intn(n int) int {
	called := r.called
	r.called++
	return r.values[called]
}
