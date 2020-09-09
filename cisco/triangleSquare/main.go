package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func main() {

	aSquare := square{
		sideLength: 10,
	}

	aTriange := triangle{
		base:   10,
		height: 10,
	}

	printArea(aTriange)
	printArea(aSquare)
}

func (square) getArea(s square) float64 {
	area := s.sideLength * s.sideLength
	return area
}

func (triangle) getArea(t triangle) float64 {
	area := 0.5 * t.base * t.height
	return area
}

func printArea(s shape) {
	fmt.Println("The area is:", s.getArea())
}

/*
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

*/
