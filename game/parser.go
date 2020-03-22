package game

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func openJSON(path string) []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	return jsonBytes
}

// CardsParser : Parser the decks database and returns a collection of them
func CardsParser() Deck {
	jsonBytes := openJSON("./game/cards.json")

	d := Deck{}

	// add unique ID to card
	for i, val := range d.Cards {
		d.Cards[i].ID = i
	}

	json.Unmarshal(jsonBytes, &d)

	return d
}
