package main

import "fmt"

// structs are custom types
type person struct {
    firstName string
    lastName string
}

func main() {
    // attribute creation/reference with dot-notation
    var alex person
    alex.firstName = "alex"
    alex.lastName = "anderson"
    fmt.Println(alex.lastName)

    // ZEN
    // explicit better than implicit
    jenkins := person{firstName: "DEEPLOY", lastName: "JENKINS!"}
    // printf debugging
    fmt.Printf("%+v\n", jenkins)

    // this looks like crap, and apparently is common
    // attributes definied with positional instead of keyword args
    jimmy := person{"jimmy", "woods"}
    fmt.Println(jimmy.firstName)
}
