package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Error("Expected length of new deck to be 52, but got", len(d))
	}

	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	var card string

	for i := range cardSuits {
		for j := range cardValues {
			card = cardValues[j] + " of " + cardSuits[i]

			if d[len(cardValues)*i+j] != card {
				t.Error("Expected card at index", x, "to be", card, "but got", d[len(cardValues)*i+j])
			}
		}
	}
}
