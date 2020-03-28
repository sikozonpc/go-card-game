package game

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// DeckMaxLen : Maximum length of deck
const DeckMaxLen = 30

// Card : Card
type Card struct {
	ID     string `json:"id"`
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
	health      int
	mana        int
	Graveyard   []Card
	Hand        []Card
	Battlefield []Card
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
		PlayerOne: Player{30, 0, []Card{}, hand1, []Card{}},
		PlayerTwo: Player{30, 0, []Card{}, hand2, []Card{}},
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
	cards := make([]Card, DeckMaxLen)
	deckCollection := CardsParser()

	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	deckCollection.Shuffle()

	//FIXME: Some cards are getting duplicates cause of this, since the deck requires more
	// cards then what the colletion has
	for i := range cards {
		randint := rnd.Intn(len(deckCollection.Cards))
		cards[i] = deckCollection.Cards[randint]
	}

	return Deck{cards}
}

// Attack : A card attacks another card
func (p *Player) Attack(attacker Card, defender Card, defenderPlayer Player) error {
	if !isCardInBattlefield(attacker, *p) {
		return errors.New("card is not in the battlefield and it's trying to attack")
	}
	if !isCardInBattlefield(defender, defenderPlayer) {
		return errors.New("card is not in the battlefield and it's trying to attack")
	}

	log.Printf("Card %v ATTACKED %v", attacker, defender)
	// apply the attac damage
	defender.Health -= attacker.Damage
	// recoil damage
	attacker.Health -= defender.Damage

	if attacker.Health <= 0 {
		killCard(*p, attacker)
	}
	if defender.Health <= 0 {
		killCard(defenderPlayer, defender)
	}

	log.Printf("Resulted from ATTACK %v ATTACKED %v", attacker, defender)

	return nil
}

// MoveCardToBattlefield : Moves a card to the battlefield
func (p *Player) MoveCardToBattlefield(c Card) error {
	if len(p.Hand) == 0 {
		return errors.New("player hand has no cards to be moved")
	}

	p.Battlefield = append(p.Battlefield, c)

	// remove from hand
	for i, card := range p.Hand {
		if card.ID == c.ID {
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			return nil
		}
	}

	return errors.New("could not find tje card on the battlefield")
}

func killCard(p Player, c Card) {
	// find and remove from battlefield
	for i, card := range p.Battlefield {
		if card.ID == c.ID {
			p.Battlefield = append(p.Battlefield[:i], p.Battlefield[i+1:]...)
		}
	}

	// add to player graveyard
	p.Graveyard = append(p.Graveyard, c)

	// TODO: Inflict damage to enemy player

	fmt.Printf("%v has been killed\n", c)
}

func isCardInBattlefield(c Card, p Player) bool {
	for _, card := range p.Battlefield {
		if card.ID == c.ID {
			return true
		}
	}

	return false
}
