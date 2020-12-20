package main

// Create a new type of 'deck'
// which is a slice of string

type deck []string

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}
