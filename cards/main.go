package main 

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	hand, remaining := deal(cards, 5)
	hand.print()
	remaining.print()

	fmt.Println(hand.toString())

	hand.saveToFile("cards/deck.txt")

	deck2 := newDeckFromFile("cards/deck.txt")
	deck2.print()

	deck2.shuffleDeck()
	deck2.print()
}