package main

func main() {
	cards := deck{}
	cards = newDeck()
	slicedDeck, remainingDeck := cards.deal(5)
	slicedDeck.printDeckSize()
	cards.printDeckSize()
	remainingDeck.saveToFile("my_cards")
	nextDeck := newDeckFromFile("my_cards")
	nextDeck.printDeckSize()
}
