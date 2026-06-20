package deck

import (
	"fmt"
	"sort"
)

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

func CreateDeck(opts ...func([]Card) []Card) []Card {

	var cards []Card

	for _, suit := range suits {
		for rank := minCard; rank <= maxCard; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards

}

func DefaultDeckSort(cards []Card) []Card {

	sort.Slice(cards, Less(cards))
	return cards

}

func Less(cards []Card) func(i, j int) bool {
	return func(i int, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxCard) + int(c.Rank)
}
