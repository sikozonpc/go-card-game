package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sikozonpc/go-card-game/server/game"
)

// Board : Local game board data
var Board game.Board

// flag to tell if there is a game running
var gameOngoing = false

// clients poll
var clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/newgame", newGameHandler)
	r.HandleFunc("/game", currenGameHandler)
	r.HandleFunc("/ws", wsEndpoint)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	// Start listening for incoming chat messages
	go handleMessages()

	log.Println("Server is running on :8083")

	log.Fatalln(http.ListenAndServe(":8083", c.Handler(r)))
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[API]: New game started from: %v", r.RemoteAddr)

	Board = game.PopulateBoard()
	gameOngoing = true

	response := clientResponse{Action: "@NEWGAME", Data: Board}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalln(err)
		return
	}

	w.Write(jsonData)
}

func currenGameHandler(w http.ResponseWriter, r *http.Request) {
	if gameOngoing {
		// Serve the board if there is one
		log.Println("Serving on going game")
		session := clientResponse{Action: "@SYNC", Data: Board}

		jsonData, err := json.Marshal(session)
		if err != nil {
			log.Fatalln(err)
			return
		}

		w.Write(jsonData)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Client Connected %v \n", r.RemoteAddr)

	wsReader(ws)

}

func wsReader(conn *websocket.Conn) {
	defer func() {
		fmt.Println("Closing connection...")
		//todo: remove from users lists
		conn.Close()
	}()

	// Register our new client
	clients[conn] = true

	fmt.Printf("%+v", clients)

	// the for is important to keep the connection open
	for {

		//TODO: In the future we'll use pool conns to control the clients

		var msg clientData

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error", err)
			delete(clients, conn)
			break
		}
		fmt.Println(msg)

		data := handleData(msg, conn)

		if data != nil {
			conn.WriteJSON(data)
		}

		// Send the newly received message to the broadcast channel
		broadcast <- data
	}
}

func handleData(data clientData, conn *websocket.Conn) interface{} {
	fmt.Printf("GOT MESASGE FROM CLIENT:  %v", data)

	switch data.Action {
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

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
