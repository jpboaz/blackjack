package main

import (
	"deck"
	"fmt"
)

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}
}
