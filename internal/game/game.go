package game

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/midgarco/dice_game/internal/die"
	"github.com/midgarco/dice_game/internal/player"
)

// Game ...
type Game struct {
	Players       []*player.Player `json:"players"`
	Dice          []*die.Die       `json:"dice"`
	CurrentPlayer *player.Player   `json:"current_player"`
	MaxScore      int
	OpenScore     int
}

// Start the game
func (g *Game) Start() {
	if g.OpenScore == 0 {
		g.OpenScore = 650
	}
	g.Dice = make([]*die.Die, 5)
	for d := range g.Dice {
		g.Dice[d] = die.NewSixSidedDie()
	}
	g.NextTurn()
}

// NextTurn ...
func (g *Game) NextTurn() {
	g.CurrentPlayer, g.Players = g.Players[0], g.Players[1:]

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Player %s [%d], it's your turn: [R]oll\n", g.CurrentPlayer.Name, g.CurrentPlayer.Score)
	if scanner.Scan() {
		if strings.ToLower(scanner.Text()) != "r" {
			g.EndTurn()
			return
		}
	}

	g.CurrentPlayer.Turn.New()
	for {
		if g.CurrentPlayer.Turn.RemainingDie == 0 {
			g.CurrentPlayer.Turn.RemainingDie = len(g.Dice)
		}

		diceRoll := g.CurrentPlayer.Turn.Roll(g.Dice[:g.CurrentPlayer.Turn.RemainingDie])
		fmt.Printf("%s's Roll >> ", g.CurrentPlayer.Name)
		for i := range diceRoll {
			fmt.Printf("[%d]:%d ", i, diceRoll[i].Value)
		}
		runningTally := die.Tally(diceRoll)
		fmt.Printf(">> %d\n", runningTally)

		if runningTally == 0 {
			break // turn is over!
		}

		// prompt to bank
		fmt.Println("Select which die to bank: ")
		var bank string
		fmt.Scanln(&bank)
		banked := g.Bank(diceRoll, bank)
		err := g.CurrentPlayer.Turn.Bank(banked)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("Banked score: ", g.CurrentPlayer.Turn.Banked)

		// prompt to save their score
		fmt.Println("remaining die", g.CurrentPlayer.Turn.RemainingDie)
		if (g.CurrentPlayer.Score >= g.OpenScore || runningTally >= g.OpenScore || g.CurrentPlayer.Turn.Banked >= g.OpenScore) && g.CurrentPlayer.Turn.RemainingDie > 0 && g.CurrentPlayer.Turn.Banked > 0 {
			fmt.Println("Do you want to [S]ave?")
			if scanner.Scan() {
				if strings.ToLower(scanner.Text()) == "s" {
					g.CurrentPlayer.Save()

					// winner winner, chicken dinner
					if g.IsWinner(g.CurrentPlayer) {
						g.End()
						return
					}

					break
				}
			}
		}
	}

	fmt.Printf("Player %s, your score is: %d\n", g.CurrentPlayer.Name, g.CurrentPlayer.Score)

	g.EndTurn()
	return
}

// Bank ...
func (g *Game) Bank(dice []*die.Die, bank string) []*die.Die {
	var banked []*die.Die
	for _, x := range bank {
		if idx, err := strconv.Atoi(string(x)); err == nil {
			banked = append(banked, dice[idx])
		}
	}
	return banked
}

// EndTurn ...
func (g *Game) EndTurn() {
	g.Players = append(g.Players, g.CurrentPlayer)
	g.CurrentPlayer = &player.Player{}
	g.NextTurn()
}

// End the game
func (g *Game) End() {
	g.Players = append(g.Players, g.CurrentPlayer)
	g.CurrentPlayer = &player.Player{}
	fmt.Printf("\n\nGame Over!\n")
	fmt.Printf("\n\nFinal Score\n")
	sort.Slice(g.Players, func(a, b int) bool {
		return g.Players[a].Score > g.Players[b].Score
	})
	for _, player := range g.Players {
		fmt.Printf("%s: %d ", player.Name, player.Score)
		if g.IsWinner(player) {
			fmt.Print("**winner**")
		}
		fmt.Print("\n")
	}
}

// IsWinner ...
func (g Game) IsWinner(p *player.Player) bool {
	return p.Score >= g.MaxScore
}
