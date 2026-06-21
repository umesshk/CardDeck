package deck

import (
	"fmt"
	"testing"
)

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

func TestDefaulDeckSort(t *testing.T) {
	cards := CreateDeck(DefaultDeckSort)
	exp := Card{Rank: Ace, Suit: Spade}

	if cards[0] != exp {
		t.Errorf("Cards Not sorted  expected %v got %v", exp, cards[0])
	}
}

func TestShuffle(t *testing.T) {
	cards := CreateDeck()
	ShuffledCards := CreateDeck(Shuffle)

	if cards[0] == ShuffledCards[0] {
		t.Error("Cards Not Shuffled ")
	}
}

func TestJoker(t *testing.T) {
	cards := CreateDeck(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Errorf("Wrong Number of Jokers appended wanted %d , got %d ", 3, count)
	}

}
