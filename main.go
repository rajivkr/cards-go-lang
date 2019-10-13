package main

import "fmt"

func main() {
	cards := []string{
		"Four of Diamonds",
		newCard(),
	}
	cards = append(cards, "Six of Diamonds")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
