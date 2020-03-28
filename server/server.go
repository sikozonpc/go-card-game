package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rs/cors"
	"github.com/sikozonpc/go-card-game/server/api"
	"github.com/sikozonpc/go-card-game/server/game"
	"golang.org/x/net/websocket"
)

// Board : Local game board data
var Board game.Board

type clientData struct {
	Action string
	Data   struct {
		Card game.Card
		To   string
	}
}

type clientResponse struct {
	Action string
	Data   interface{}
}

func main() {
	go startRestAPI()
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

func startRestAPI() {
	s := &api.RestServer{}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	http.Handle("/api", s)
	err := http.ListenAndServe(":8083", c.Handler(s))

	log.Println("[API]: Listening API on localhost:8083")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func startWebSocket() {
	http.Handle("/", websocket.Handler(handleWebSocket))
	err := http.ListenAndServe(":8082", nil)

	fmt.Println("[WS]: Listening WS on localhost:8082")

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

	switch data.Action {
	case "@NEW-GAME":
		{
			// Create a new battle session
			Board = game.PopulateBoard()

			return clientResponse{"@NEW-GAME", Board}
		}
	case "@MOVE-CARD-TO-BATTLEFIELD":
		{

			var playedCard interface{} = data.Data.Card
			c, ok := playedCard.(game.Card)
			if ok {
				Board.PlayerOne.MoveCardToBattlefield(c)
				return clientResponse{"@MOVE-CARD-TO-BATTLEFIELD", Board}
			}

			log.Println("Failed to convert playedCard to card")
			return clientResponse{"@ERROR", "Failed to convert playerCard to card"}
		}
	}

	return nil
}

func handleWebSocket(ws *websocket.Conn) {
	handleConnection(ws)
}
