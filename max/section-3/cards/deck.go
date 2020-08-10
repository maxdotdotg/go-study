package main

import "fmt"
// create a new type, `deck`
// which is a slice of strings
type deck []string

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
