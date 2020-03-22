package main

import (
	"github.com/sikozonpc/go-card-game/game"
)

func main() {

	b := game.PopulateBoard()

	playerOneHand := b.PlayerOne.Hand
	playerTwoHand := b.PlayerTwo.Hand
}
