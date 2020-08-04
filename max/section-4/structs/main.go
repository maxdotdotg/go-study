package main

import "fmt"

// structs are custom types, LGTM
type person struct {
    firstName string
    lastName string
}


func main() {
    // we can intantiate a struct using positional args
    // and this will break if anything in the struct changes
    // not great
    alex := person{"Mr.", "Anderson"}
    fmt.Println("variable alex", alex)

    // we can make the properties explicit tho
    jenkins := person{firstName:"DEEEEEPLOY", lastName: "JENKINS"}
    fmt.Println("variable jenkins", jenkins)

    // we're declaring a variable jimmy of type `person`
    // and the variable is initialized with null values
    // "" in the case of string
    var jimmy person
    fmt.Println(jimmy) // prints `{ }`

    // printf debuggin, maybe `%+v` for verbose?
    // add a \n for readability
    fmt.Printf("%+v\n", jimmy)

    // update values in a struct
    jimmy.firstName = "jimmy"
    jimmy.lastName = "woods"
    fmt.Printf("%+v", jimmy)
}
