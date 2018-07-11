package game

import (
	"testing"

	"github.com/midgarco/dice_game/internal/die"
)

type Test struct {
	Dice          []*die.Die
	Bank          string
	ExpectedScore int
}

var tests = []Test{
	{Dice: []*die.Die{&die.Die{Value: 1}, &die.Die{Value: 2}, &die.Die{Value: 5}, &die.Die{Value: 5}, &die.Die{Value: 5}}, Bank: "0234", ExpectedScore: 600},
}

// [0]:1 [1]:2 [2]:5 [3]:5 [4]:5 >> 600
// Select which die to bank
// 0234
// Banked score:  200
func TestBank(t *testing.T) {
	for x := range tests {
		test := tests[x]
		g := &Game{}
		banked := g.Bank(test.Dice, test.Bank)
		score := die.Tally(banked)
		if score != test.ExpectedScore {
			t.Errorf("Expected %d, got %d", test.ExpectedScore, score)
		}
	}
}
