package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/midgarco/dice_game/internal/game"
	"github.com/midgarco/dice_game/internal/player"
)

func main() {
	rand.Seed(time.Now().Unix())

	game := &game.Game{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many players? ")
	playerCount, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	numPlayers, err := strconv.Atoi(playerCount)
	if err != nil {
		panic(err)
	}
	game.Players = make([]*player.Player, numPlayers)

	fmt.Println(game)
}
