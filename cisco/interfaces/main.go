package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// logic
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// logc
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
