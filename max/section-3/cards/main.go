package main

func main() {
	// long variable initialization form
	//var card string = "Ace of Spades"

	// short variable initialization form
	// let the compiler figure it out
	//card := "Ace of Spades"

	// colon-equals only used when initializing a variable
	// re-assigning a variable is only equals
	//card = "Five of Diamonds"
	//fmt.Println(card)

	// variables can't be assigned outside of a function?
	// yeah, idk, makes me curious about global vars and such
	//deckSize = 52
	//fmt.Println(deckSize)

	// use the newCard function to instantiate a
	// you guessed it!
	// new card
	//thirdCard := newCard()
	//fmt.Println(thirdCard)

	// let's make a slice
	// again, that's an array of unfixed size
	// type declaration required
	// a slice of type string
	// both strings and functions that return strings are valid
	// cards := []string{"Ace of Spades", newCard()}

	// we made a type called deck, so lets instatiate that instead
	//cards := deck{"Ace of Spades", newCard()}

	// let's add an item to the slice
	// a new slice is created using the current content of `cards` plus
	// the new item, and a new slice is returned
	// that new since is assigned to the variable `cards`
	// this makes me feel not great, and I am not sure why
	// maybe it's the mutable part? IDK
	//cards = append(cards, "Six of Clubs")

	// let's iterate
	// for index and card in the slice/list/thing that is cards
	// print the index and the data at that index
	// the colon-equals is used because the variables i and card
	// go out-of-scope as the loop closes
	//for i, card := range cards {
	//    fmt.Println(i, card)
	//}
	// in python:
	// for i in cards:
	//    print(cards)

	// make a new deck, this time using the custom type
	cards := newDeck()

	// we made a function called print in deck.go that can be used
	// on type `deck` so let's use that instead
	//cards.print()

	// I am a bit confused by this
	// the index returned, not the content of `cards`
	// this is probably my defaulting to python
	//for card := range cards {
	//    fmt.Println(card)
	//}
	// the loop above does this:
	// for i in cards:
	//    print(cards.index(i))

	// deal a hand
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()

	// write the deck to stdout as one long string
	//fmt.Println(cards.toString())

	// save the deck to the file system
	cards.saveToFile("test-deck")

	new_deck := newDeckFromFile("test-deck")
	new_deck.shuffleDeck()
	new_deck.print()
}

// function signatures require return type declarations
//func newCard() string {
//    return "Five of Hearts"
//}
