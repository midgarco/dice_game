package player

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/midgarco/dice_game/internal/die"
)

// Player ...
type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Turn  *Turn  `json:"turn"`
}

// Turn ...
type Turn struct {
	Banked       int `json:"banked"`
	RemainingDie int `json:"remaining_die"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NewPlayer sets up the player details
func NewPlayer(num int) *Player {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Player %d Name: \n", num)
	player := &Player{}
	player.Turn = &Turn{}

	// if scanner.Scan() {
	// 	player.Name = scanner.Text()
	// }

	_ = scanner
	player.Name = string(letterBytes[rand.Int63()%int64(len(letterBytes))])
	return player
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
		t.Banked += dice[d].Value
	}
	if t.RemainingDie == 0 {
		t.RemainingDie = 5
	}
}

// Reset the turn
func (t *Turn) Reset() {
	t.Banked = 0
	t.RemainingDie = 5
}
