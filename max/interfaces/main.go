package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main(){
}

// if you're not gonna use the receiver's object/variable/thing
// itself, you don't have to include it when writing the receiver
// just the type
func (englishBot) getGreeting() string {
    // some specific logic
    // also remmeber you don't need fmt to return a string
    return "Hello there"
}

func (spanishBot) getGreeting() string {
    // some other specific logic
    return "Hola"
}

