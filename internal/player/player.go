package player

import (
	"github.com/midgarco/dice_game/internal/die"
)

// Player ...
type Player struct {
	Name        string `json:"name"`
	Score       int    `json:"score"`
	CurrentTurn *Turn  `json:"current_turn"`
}

// Turn ...
type Turn struct {
	Score        int `json:"score"`
	RemainingDie int `json:"remaining_die"`
}

// New turn
func (t *Turn) New() {
	t.Reset()
}

// Roll the dice
func (t Turn) Roll(dice []*die.Die) []*die.Die {
	var result []*die.Die
	for d := range dice {
		dice[d].Roll()
		result = append(result, dice[d])
	}
	return result
}

// Bank the current score
func (t *Turn) Bank(dice []*die.Die, saves ...int) {
	for d := range dice {
		t.RemainingDie--
		t.Score += dice[d].Value
	}
	if t.RemainingDie == 0 {
		t.RemainingDie = 5
	}
}

// Reset the turn
func (t *Turn) Reset() {
	t.Score = 0
	t.RemainingDie = 5
}
