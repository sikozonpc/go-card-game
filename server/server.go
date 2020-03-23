package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/sikozonpc/go-card-game/game"
)

func main() {
	go server()

	//let the server goroutine run forever
	var input string
	fmt.Scanln(&input)
}

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer func() {
		listener.Close()
		fmt.Println("Listener closed")
	}()

	fmt.Println("Server started successfully!")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Handling new connection...")

	// Close connection when this function ends
	defer func() {
		fmt.Println("Closing connection...")
		conn.Close()
	}()

	// creates a decoder that reads directly from the socket
	d := json.NewDecoder(conn)

	var msg game.Card

	err := d.Decode(&msg)

	fmt.Println(msg, err)
}
