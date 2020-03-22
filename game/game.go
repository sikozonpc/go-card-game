package game

import (
	"fmt"
	"math/rand"
	"time"
)

// DeckNaxLen : Maximum length of deck
const DeckNaxLen = 30

// Card : Card
type Card struct {
	Name   string  `json:"name"`
	Health float32 `json:"health"`
	Damage float32 `json:"damage"`
}

// Deck : Collection of cards
type Deck struct {
	Cards []Card
}

func Init() {
	d := CreateDeck()
	d.Shuffle()
	fmt.Println(len(d.Cards), d)
}

// Shuffle : Shuffles the deck
func (d Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	for i := range d.Cards {
		randint := rnd.Intn(len(d.Cards))
		d.Cards[i] = d.Cards[randint]
	}
}

// CreateDeck : Randomly selects from the collection of cards and creates a Deck
func CreateDeck() Deck {
	cards := make([]Card, DeckNaxLen)
	deckCollection := CardsParser()

	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	for i := range cards {
		randint := rnd.Intn(len(deckCollection.Cards))
		cards[i] = deckCollection.Cards[randint]
	}

	return Deck{cards}
}
