package game

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
	pwd, _ := os.Getwd()
	jsonBytes := openJSON(pwd + "/game/cards.json")

	d := Deck{}

	json.Unmarshal(jsonBytes, &d)

	// add unique uuid to a card
	for i := range d.Cards {
		d.Cards[i].ID = strconv.Itoa(i)
	}

	return d
}
