package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cards.populate()
	return cards
}

func newEmptyDeck() deck {
	cards := deck{}
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

func (d deck) printToFile(filename string) error {

	str := strings.Join(d, "\n")
	return ioutil.WriteFile(filename, []byte(str), 0777)
}

func (d deck) printUsingOS() {
	str := strings.Join(d, "\n")
	file, err := os.Open("deck.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(str))
}

func createDeckFromFile(filename string) (deck, error) {
	// str, err := ioutil.ReadFile("mydeck.txt")
	// newDeck := deck(strings.Split(string(str), "\n"))
	// return newDeck, err
	cards := newEmptyDeck()
	err := cards.populateFromFile(filename)
	return cards, err
}

func (d *deck) populateFromFile(filename string) error {
	str, err := ioutil.ReadFile("mydeck.txt")
	if err == nil {
		cards := strings.Split(string(str), "\n")
		for _, card := range cards {
			*d = append(*d, card)
		}
	}
	return err
}

func (d deck) shuffle() {
	//	var newPosition int
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
