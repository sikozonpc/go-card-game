package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/sikozonpc/go-card-game/game"
	"golang.org/x/net/websocket"
)

type clientData struct {
	Action string
	Data   interface{}
}

func main() {
	//go startTCP()
	go startWebSocket()

	//let the server goroutine run forever
	var input string
	fmt.Scanln(&input)
}

func startTCP() {
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

func startWebSocket() {
	http.Handle("/", websocket.Handler(handleWebSocket))
	err := http.ListenAndServe(":8082", nil)

	fmt.Println("Listening WS on " + "localhost:8082")

	if err != nil {
		fmt.Println(err)
		panic(err)
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

	var msg clientData

	err := d.Decode(&msg)

	fmt.Println(msg, err)

	data := handleData(msg, conn)
	if data != nil {
		jsonMsg, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
			return
		}

		conn.Write(jsonMsg)
	}
}

func handleData(data clientData, conn net.Conn) interface{} {
	fmt.Printf("GOT MESASGE FROM CLIENT:\n %v", data)

	if data.Action == "@NEW-GAME" {
		// Create a new battle session
		board := game.PopulateBoard()

		return clientData{"@NEW-GAME", board}
	}

	return nil
}

func handleWebSocket(ws *websocket.Conn) {
	handleConnection(ws)
}
