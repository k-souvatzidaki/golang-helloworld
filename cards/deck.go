package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// constructor
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return two decks: one for the player's cards and one for the rest
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert deck to string representation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save deck to file 
func (d deck) saveToFile(filename string) error { // the WriteFile function returns an error
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// read deck from file 
func newDeckFromFile(filename string) deck{
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(content), ","))
}

// randomize cards order
func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano()) // the seed
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) -1)

		// swap
		d[i], d[newPos] = d[newPos], d[i]
	}
}