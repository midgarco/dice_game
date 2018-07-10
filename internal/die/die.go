package die

import (
	"math/rand"
	"sort"
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
	scored := []*Die{}
	remain := []*Die{}

	sort.Slice(dice, func(a, b int) bool {
		return dice[a].Value < dice[b].Value
	})

	if len(dice) == 5 {
		// run !
		if dice[0].Value < dice[1].Value && dice[1].Value < dice[2].Value && dice[2].Value < dice[3].Value && dice[3].Value < dice[4].Value && (dice[4].Value-dice[0].Value) < 5 {
			scored = dice
			total += 1500
			dice = remain
			return total
		}

		// 5 of a kind !
		if dice[0].Value == dice[1].Value && dice[1].Value == dice[2].Value && dice[2].Value == dice[3].Value && dice[3].Value == dice[4].Value {
			scored = dice
			if dice[0].Value == 1 {
				total += 10000
			} else {
				total += dice[0].Value * 1000
			}
			dice = remain
			return total
		}
	}

	// 3 of a kind
	if len(dice) >= 3 && dice[0].Value == dice[1].Value && dice[1].Value == dice[2].Value {
		scored = append(scored, dice[:2]...)
		remain = dice[3:]
		if dice[0].Value == 1 {
			total += 1000
		} else {
			total += dice[0].Value * 100
		}
		dice = remain
	} else if len(dice) >= 4 && dice[1].Value == dice[2].Value && dice[2].Value == dice[3].Value {
		scored = append(scored, dice[1:3]...)
		remain = append(remain, dice[0])
		if len(dice) == 5 {
			remain = append(remain, dice[4])
		}
		total += dice[1].Value * 100
		dice = remain
	} else if len(dice) == 5 && dice[2].Value == dice[3].Value && dice[3].Value == dice[4].Value {
		scored = append(scored, dice[2:4]...)
		remain = dice[:2]
		total += dice[2].Value * 100
		dice = remain
	}

	// all the rest
	for d := range dice {
		if dice[d].Value == 1 {
			total += 100
			scored = append(scored, dice[d])
		} else if dice[d].Value == 5 {
			total += 50
			scored = append(scored, dice[d])
		} else {
			remain = append(remain, dice[d])
		}
	}

	return total
}
