package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	// convert the deck to slice of strings `[]string(d)`
	// convert slice of strings to one big string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// using the WriteFile function from the ioutil package
	// write the file after converting `deck` to string to bytes
	// file permissions are part of the signature for WriteFile

	// FIXME: I have no real idea why we're returning the error type
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// read from file, and capture both the byte slice and the error
	// if error's not empty, break and exit 1
	// lots of this pattern in go, if error exists, then exit
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
