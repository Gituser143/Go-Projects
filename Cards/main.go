package main

import "fmt"

func main() {

	var cards = newDeck()
	hand, cards := deal(cards, 4)

	fmt.Println("Cards:")
	cards.print()

	fmt.Println("\nHand:")
	hand.print()
}
