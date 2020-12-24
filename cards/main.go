package main

func main() {
	cards := newDeck()

	cards.shuffle()

	hand, remaining := deal(cards, 3)

	hand.saveToFile("my_cards")

	deck := readFromFile("my_cards")
	deck.print()

	remaining.print()

}

func newCard() string {
	return "Five of Diamonds"
}
