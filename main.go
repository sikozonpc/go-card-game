package main

import (
	"fmt"

	"github.com/sikozonpc/go-card-game/game"
)

func main() {

	b := game.PopulateBoard()
	fmt.Print(b.PlayerOne.Hand)

}
