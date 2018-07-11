package player

import (
	"testing"

	"github.com/midgarco/dice_game/internal/die"
)

type Test struct {
	Dice          []*die.Die
	ExpectedScore int
}

var tests = []Test{
	{Dice: []*die.Die{&die.Die{Value: 1}, &die.Die{Value: 5}, &die.Die{Value: 5}, &die.Die{Value: 5}}, ExpectedScore: 600},
}

func TestBank(t *testing.T) {
	for x := range tests {
		test := tests[x]
		turn := &Turn{RemainingDie: len(test.Dice)}
		turn.Bank(test.Dice)
		if turn.Banked != test.ExpectedScore {
			t.Errorf("Expected %d, got %d", test.ExpectedScore, turn.Banked)
		}
	}
}
