package main

import (
	"log"
	"os"
	"testing"
)

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
				t.Error("Expected card at index", len(cardValues)*i+j, "to be", card, "but got", d[len(cardValues)*i+j])
			}
		}
	}
}

func TestCreateDeckFromFileAndPrintToFile(t *testing.T) {
	os.Remove("_testFile")
	d := newDeck()
	err := d.printToFile("_testFile")

	if err != nil {
		t.Error("Error while saving to _testFile")
		log.Fatal(err)
	}

	deckFromFile, err := createDeckFromFile("_testFile")

	if err != nil {
		t.Error("Error while reading from _testFile")
		log.Fatal(err)
	}

	if len(deckFromFile) != 52 {
		t.Error("Expected length of deck from file to be 52, but got", len(deckFromFile))
	}

	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	var card string

	for i := range cardSuits {
		for j := range cardValues {
			card = cardValues[j] + " of " + cardSuits[i]

			if deckFromFile[len(cardValues)*i+j] != card {
				t.Error("Expected card at index", len(cardValues)*i+j, "to be", card, "but got", deckFromFile[len(cardValues)*i+j])
			}
		}
	}
	os.Remove("_testFile")

}
