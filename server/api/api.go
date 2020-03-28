package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sikozonpc/go-card-game/server/game"
)

// RestServer : Rest API server
type RestServer struct{}

func (s *RestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[API]: Connection from \n %v", r.RemoteAddr)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	Board := game.PopulateBoard()

	jsonData, err := json.Marshal(Board)
	if err != nil {
		log.Fatalln(err)
		return
	}

	w.Write(jsonData)
}
