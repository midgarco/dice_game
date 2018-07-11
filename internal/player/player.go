package player

import (
	"bufio"
	"fmt"
	"os"
	"sort"

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

	if scanner.Scan() {
		player.Name = scanner.Text()
	}

	// _ = scanner
	// player.Name = string(letterBytes[rand.Int63()%int64(len(letterBytes))])

	return player
}

// Save the banked score and reset the turn
func (p *Player) Save() {
	p.Score += p.Turn.Banked
	p.Turn.Reset()
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
	sort.Slice(result, func(a, b int) bool {
		return result[a].Value < result[b].Value
	})
	return result
}

// Bank the current score
func (t *Turn) Bank(dice []*die.Die) error {
	if len(dice) > t.RemainingDie {
		return fmt.Errorf("to many dice relative to how many we expected. got %d, expected %d", len(dice), t.RemainingDie)
	}

	t.RemainingDie -= len(dice)
	t.Banked += die.Tally(dice)

	return nil
}

// Reset the turn
func (t *Turn) Reset() {
	t.Banked = 0
	t.RemainingDie = 5
}
