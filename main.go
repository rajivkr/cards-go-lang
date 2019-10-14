package main

func main() {
	cards := deck{}
	cards = newDeck()
	//cards.printDeck()
	slicedDeck, remainingDeck := deal(cards, 5)
	slicedDeck.printDeck()
	remainingDeck.printDeck()

}
