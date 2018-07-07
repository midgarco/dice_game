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
	total := 0
	present := map[int]int{}

	for d := range dice {
		if _, ok := present[dice[d].Value]; !ok {
			present[dice[d].Value] = 0
		}
		present[dice[d].Value]++
	}

	// 5 of a kind
	if len(dice) == 5 && len(present) == 1 {
		if dice[0].Value == 1 {
			total += 10000
		} else {
			total += dice[0].Value * 1000
		}
	}

	fmt.Printf("\nnumber of values: %d ", len(present))
	for d, count := range present {
		fmt.Printf("\n%d: %d ", d, count)
	}

	return total
}
