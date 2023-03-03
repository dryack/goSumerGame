package gameplay

import (
	"math/rand"
)

// stochasticRoundDivide accepts a dividend and divisor, both integers, and
// returns a result integer rounded according to a probability which is
// proportional to the decimal fraction part of the result. For example, 200/7
// will round up approximately 57.09% of the time.
func stochasticRoundDivide(dividend, divisor int) int {
	result := float64(dividend) / float64(divisor)
	fraction := result - float64(int(result))
	if rand.Float64() < fraction {
		return int(result) + 1
	} else {
		return int(result)
	}
}
