package main

import (
	"log"
)

func main() {

	var cards = newDeck()
	err := cards.printToFile("mydeck.txt")
	if err != nil {
		log.Fatal(err)
	}

	newCards, err := createDeckFromFile("mydeck.txt")

	if err != nil {
		log.Fatal(err)
	}
	newCards.print()
}
