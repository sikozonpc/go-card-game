package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/sikozonpc/go-card-game/game"
)

func main() {
	//TODO: Make these variable global

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer func() {
		conn.Close()
		fmt.Println("Connection closed")
	}()

	// send a message
	cardData := game.Card{
		ID:     "666",
		Damage: 5,
		Health: 2,
		Name:   "Bidoof",
	}

	//encoder := json.NewEncoder(conn)
	//decoder := json.NewDecoder(conn)

	jsonData, err := json.Marshal(cardData)
	if err != nil {
		log.Fatalln(err)
	}

	conn.Write(jsonData)

}
