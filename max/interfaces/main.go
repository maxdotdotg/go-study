package main

import "fmt"

type bot interface{
    // defining shared functions? sorta...
    // if another type/struct has this function signature
    // then it's ALSO this type
    // ANOTHER WAY
    // if a type/struct satisfies the interface, 
    // it has access to the inerface's functions

    // func name and return type
    getGreeting() string
    // because the englishBot and spanishBot type have a getGreeting()
    // function associated with them, they're ALSO of type bot 
    // and can use printGreeting()
}

type englishBot struct{}
type spanishBot struct{}

func main(){
    eng := englishBot{}
    printGreeting(eng)

    span := spanishBot{}
    printGreeting(span)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

// if you're not gonna use the receiver's object/variable/thing
// itself, you don't have to include it when writing the receiver
// just the type
// these receivers can have the same name becaues they're acting
// on different types?
func (englishBot) getGreeting() string {
    // some specific logic
    // also remmeber you don't need fmt to return a string
    return "Hello there"
}

func (spanishBot) getGreeting() string {
    // some other specific logic
    return "Hola"
}

/*
// no overloading with functions of the same name
func printGreeting(eb englishBot) {
    fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
    fmt.Println(sb.getGreeting())
}
*/
