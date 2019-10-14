package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamond", "Clever"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suitName := range cardSuits {
		for _, valueName := range cardValues {
			cards = append(cards, valueName+" of "+suitName)
		}
	}
	return cards

}

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) printDeckSize() {
	count := 0
	for range d {
		count++
	}
	fmt.Println(count)
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
