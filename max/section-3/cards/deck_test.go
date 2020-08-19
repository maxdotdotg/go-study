package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	actualLength := len(d)
	if expectedLength != actualLength {
		t.Errorf("expected deck length of %d, got %v", expectedLength, actualLength)
	}

	// I don't know if this is actually a good test, but whatever
	expectedFirstCard := "Ace of Spades"
	actualFirstCard := d[0]
	if expectedFirstCard != actualFirstCard {
		t.Errorf("expected first card to be %v, got %v", expectedFirstCard, actualFirstCard)
	}

	// no mechanism for counting backwards on indices, kind of annoying
	expectedLastCard := "King of Hearts"
	actualLastCard := d[len(d)-1]
	if expectedLastCard != actualLastCard {
		t.Errorf("expected first card to be %v, got %v", expectedLastCard, actualLastCard)
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktest")

	d := newDeck()
	d.saveToFile("_decktest")

	loadedDeck := newDeckFromFile("_decktest")

	/*
	    this doesn't work, and it's annoying
	   I can't compare these two
	   "slice can only be compared to nil"
	   slices can only be checked if they're empty/nil or not
	   that feels ugly but /shrug
	   if d != loadedDeck {
	       t.Errorf("expected deck to be %v, got %v", loadedDeck, d)
	   }
	*/

	/*
	   this _also_ doesn't work
	   thought comparing arrays after casting them would work
	   guess not
	   more /shrug
	   dArray := [len(d)]string(d)
	   loadedArray := [len(loadedDeck]string(loadedDeck)

	   if dArray != loadedArray {
	       t.Errorf("expected deck to be %v, got %v", loadedArray, d)
	   }
	*/

	/*
	   again, this feels like bullshit
	   I want to take an md5 of the string slice, but that's a lot
	   maybe something to look into later
	   https://mrwaggel.be/post/golang-hash-sum-and-checksum-to-string-tutorial-and-examples/
	   a bunch more /shrug
	*/
	expectedLength := len(d)
	actualLength := len(loadedDeck)
	if expectedLength != actualLength {
		t.Errorf("expected deck length of %d, got %v", expectedLength, actualLength)
	}

	os.Remove("_decktest")
}
