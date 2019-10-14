package main

import "fmt"

func main() {
	cards := deck{}
	cards = newDeck()
	slicedDeck, remainingDeck := cards.deal(5)
	remainingDeck.printDeckSize()
	cards.printDeckSize()
	fmt.Println(slicedDeck.toString())
}
