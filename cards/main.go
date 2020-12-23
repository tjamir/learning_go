package main

func main() {
	cards := newDeck()

	hand, _ := deal(cards, 3)

	hand.saveToFile("my_cards")

	deck := readFromFile("my_cards")
	deck.print()

}

func newCard() string {
	return "Five of Diamonds"
}
