package die

import "testing"

type Test struct {
	Dice          []*Die
	ExpectedScore int
}

var tests = []Test{
	{Dice: []*Die{&Die{Value: 1}, &Die{Value: 2}}, ExpectedScore: 100},
	{Dice: []*Die{&Die{Value: 1}, &Die{Value: 2}, &Die{Value: 3}, &Die{Value: 4}, &Die{Value: 5}}, ExpectedScore: 1500},
	{Dice: []*Die{&Die{Value: 2}, &Die{Value: 3}, &Die{Value: 4}, &Die{Value: 5}, &Die{Value: 6}}, ExpectedScore: 1500},
	{Dice: []*Die{&Die{Value: 1}, &Die{Value: 2}, &Die{Value: 3}, &Die{Value: 5}, &Die{Value: 6}}, ExpectedScore: 150},
	{Dice: []*Die{&Die{Value: 1}, &Die{Value: 1}, &Die{Value: 1}}, ExpectedScore: 1000},
	{Dice: []*Die{&Die{Value: 1}, &Die{Value: 1}, &Die{Value: 1}, &Die{Value: 1}, &Die{Value: 1}}, ExpectedScore: 10000},
	{Dice: []*Die{&Die{Value: 5}, &Die{Value: 5}, &Die{Value: 5}, &Die{Value: 5}, &Die{Value: 5}}, ExpectedScore: 5000},
	{Dice: []*Die{&Die{Value: 1}, &Die{Value: 5}, &Die{Value: 1}, &Die{Value: 1}, &Die{Value: 5}}, ExpectedScore: 1100},
	{Dice: []*Die{&Die{Value: 6}, &Die{Value: 6}, &Die{Value: 6}, &Die{Value: 5}, &Die{Value: 1}}, ExpectedScore: 750},
	{Dice: []*Die{&Die{Value: 6}, &Die{Value: 3}, &Die{Value: 6}, &Die{Value: 6}}, ExpectedScore: 600},
}

func TestTally(t *testing.T) {
	for x := range tests {
		test := tests[x]
		tally := Tally(test.Dice)
		if tally != test.ExpectedScore {
			t.Errorf("Expected %d, got %d", test.ExpectedScore, tally)
		}
	}
}
