package deck

import (
	"fmt"
	"testing"
)

var t *testing.T

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	// Output:
	// Ace of Hearts
}

func TestNewDeck(t *testing.T) {
	cards := CreateDeck()

	if len(cards) != 13*4 {
		t.Error("Wrong Number of Cards")
	}
}
