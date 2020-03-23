package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sikozonpc/go-card-game/game"
)

func main() {

	b := game.PopulateBoard()

	playerOneHand := b.PlayerOne.Hand
	playerTwoHand := b.PlayerTwo.Hand

	fmt.Println(playerOneHand)
	fmt.Println(playerTwoHand)

	err := b.PlayerOne.MoveCardToBattlefield(playerOneHand[0])
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	err = b.PlayerTwo.MoveCardToBattlefield(playerTwoHand[0])
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	err = b.PlayerOne.Attack(b.PlayerOne.Battlefield[0], b.PlayerTwo.Battlefield[0], b.PlayerTwo)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	fmt.Println(b.PlayerOne)
	fmt.Println(b.PlayerTwo)

}
