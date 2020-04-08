package main

import (
	"log"
)

func main() {

	var cards = newDeck()
	// hand, cards := deal(cards, 4)
	//
	// fmt.Println("Cards:")
	// cards.print()
	//
	// fmt.Println("\nHand:")
	// hand.print()
	// cards.printToFile()
	// cards.printUsingOS()
	// log.Fatal("hello")
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
