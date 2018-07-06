package game

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/midgarco/dice_game/internal/die"
	"github.com/midgarco/dice_game/internal/player"
)

// Game ...
type Game struct {
	Players       []*player.Player `json:"players"`
	Dice          []*die.Die       `json:"dice"`
	CurrentPlayer *player.Player   `json:"current_player"`
}

// Start the game
func (g *Game) Start() {
	g.Dice = make([]*die.Die, 5)
	for d := range g.Dice {
		g.Dice[d] = die.NewSixSidedDie()
	}
	g.NextTurn()
}

func (g *Game) NextTurn() {
	g.CurrentPlayer, g.Players = g.Players[0], g.Players[1:]

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Player %s, it's your turn: [R]oll\n", g.CurrentPlayer.Name)
	if scanner.Scan() {
		if strings.ToLower(scanner.Text()) != "r" {
			g.EndTurn()
			return
		}
	}

	g.CurrentPlayer.Turn.New()
	diceRoll := g.CurrentPlayer.Turn.Roll(g.Dice)

	sort.Slice(diceRoll, func(a, b int) bool {
		return diceRoll[a].Value < diceRoll[b].Value
	})

	for i := range diceRoll {
		fmt.Printf("%d ", diceRoll[i].Value)
	}
	fmt.Printf(">> %d\n", die.Tally(diceRoll))

	g.EndTurn()
	return
}

func (g *Game) EndTurn() {
	g.Players = append(g.Players, g.CurrentPlayer)
	g.CurrentPlayer = &player.Player{}
	g.NextTurn()
}
