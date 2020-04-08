package main

func main() {

	var cards = newDeck()

	cards.print()
}

func getCard() string {
	return "King of Diamonds"
}
