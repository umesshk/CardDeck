package deck

import "fmt"

//go:generate stringer -type=Suit,Rank

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank uint8

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minCard = Ace
	maxCard = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {

	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func CreateDeck() []Card {

	var cards []Card

	for _, suit := range suits {
		for rank := minCard; rank <= maxCard; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}

	return cards

}
