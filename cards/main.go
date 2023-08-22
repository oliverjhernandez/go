package main

func main() {

	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// cards.saveToFile("super_file")
	cards := newDeckFromFile("super_file")
	cards.shuffle()
	cards.print()
}
