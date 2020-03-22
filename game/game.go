package game

import (
	"log"
	"math/rand"
	"time"
)

// DeckNaxLen : Maximum length of deck
const DeckNaxLen = 30

// Card : Card
type Card struct {
	Name   string `json:"name"`
	Health int    `json:"health"`
	Damage int    `json:"damage"`
}

// Deck : Collection of cards
type Deck struct {
	Cards []Card
}

// Player : Player struct
type Player struct {
	health    int
	mana      int
	Graveyard []Card
	Hand      []Card
}

// Board : Board struct
type Board struct {
	PlayerOne Player
	PlayerTwo Player
}

// PopulateBoard : Initializes a new board and fills it with data
func PopulateBoard() Board {
	deck := CreateDeck()
	hand1 := deck.Draw(3)
	hand2 := deck.Draw(3)

	board := Board{
		PlayerOne: Player{30, 0, []Card{}, hand1},
		PlayerTwo: Player{30, 0, []Card{}, hand2},
	}

	return board
}

// Shuffle : Shuffles the deck
func (d *Deck) Shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	for i := range d.Cards {
		randint := rnd.Intn(len(d.Cards))
		d.Cards[i] = d.Cards[randint]
	}
}

// Draw : Draws `n` cards from the deck
func (d *Deck) Draw(n int) []Card {
	cards := make([]Card, n)

	for i := 0; i < n; i++ {
		// retrieve card from deck
		cards[i] = d.Cards[i]

		copy(d.Cards[i:], d.Cards[i+1:])
		d.Cards[len(d.Cards)-1] = Card{}
		d.Cards = d.Cards[:len(d.Cards)-1]
	}

	return cards
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

// Attack : A card attacks another card
func (c Card) Attack(target Card) {
	log.Printf("Card %v ATTACKED %v", c, target)

	// apply the attac damage
	target.Health -= c.Damage
	// recoil damage
	c.Health -= target.Damage

	log.Printf("Resulted from ATTACK %v ATTACKED %v", c, target)
}
