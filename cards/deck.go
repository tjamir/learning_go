package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of string

type deck []string

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func deal(deck deck, handSize int) (deck, deck) {
	hand := deck[:handSize]
	remaining := deck[handSize:]
	return hand, remaining
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func readFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log de error and return new Deck
		// Option #2 - Log de error and quit the program
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	cardsNames := string(bytes)
	return deck(strings.Split(cardsNames, ","))
}
