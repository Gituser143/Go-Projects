package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cards.populate()
	return cards
}

func (d *deck) populate() {
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	cardSuits := []string{"Diamonds", "Hearts", "Clubs", "Spades"}

	var card string

	for i := range cardSuits {
		for j := range cardValues {
			card = cardValues[j] + " of " + cardSuits[i]
			*d = append(*d, card)
		}
	}
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
