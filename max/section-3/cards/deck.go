package main

import "fmt"

// create a new type, `deck`
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// for variables that will never be used, like the index here
	// we can replace it with an underscore so the compiler
	// doesn't yell at us
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// new custom function
// this is available to any variables of type `deck`
// and it prints the content
// it gets used in the form `deck.print()`
// instead of a name, the function gets a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// perentheses used for returning multiple pieces of data
func deal(d deck, handSize int) (deck, deck) {
	// take the deck and split it
	// return two decks, handSize and the remainder
	return d[:handSize], d[handSize:]
}
