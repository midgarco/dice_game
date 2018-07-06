package die

import (
	"fmt"
	"math/rand"
)

// Die model
type Die struct {
	Sides int `json:"sides"`
	Value int `json:"value"`
}

// NewSixSidedDie quickly creates a six sided die
func NewSixSidedDie() *Die {
	return &Die{Sides: 6}
}

// Roll the die
func (d *Die) Roll() {
	d.Value = rand.Intn(d.Sides) + 1
}

// Tally the score
func Tally(dice []*Die) int {
	present := map[int]int{}

	for d := range dice {
		if _, ok := present[dice[d].Value]; !ok {
			present[dice[d].Value] = 0
		}
		present[dice[d].Value]++
	}

	for d, count := range present {
		fmt.Printf("\n%d: %d ", d, count)
	}

	total := 0
	return total
}
