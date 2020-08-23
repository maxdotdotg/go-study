package main

import "fmt"

// structs are custom types
type person struct {
    firstName string
    lastName string

    // embedded structs!
    // makes sense, just another property/type pair
    contact contactInfo
}

type contactInfo struct {
    email string
    zipCode int
}


func main() {
    // attribute creation/reference with dot-notation
    // `alex` is initialized with the Zero Value of "" type string
    // for the properties firstName and lastName
    var alex person
    fmt.Printf("%+v\n", alex)
    alex.firstName = "alex"
    alex.lastName = "anderson"
    fmt.Println(alex.lastName)

    // ZEN
    // explicit better than implicit
    // this DOESN'T break when we add another property to the struct!
    // I guess the contact property gets a zero value and that's it
    jenkins := person{firstName: "DEEPLOY", lastName: "JENKINS!"}
    // printf debugging
    fmt.Printf("%+v\n", jenkins)

    /*
    // this looks like crap, and apparently is common
    // attributes definied with positional instead of keyword args
    // also, this breaks when we add another property to the struct!
    // not sure why contact doesn't get a zero value, but it doesn't? /shrug
    jimmy := person{"jimmy", "woods"}
    fmt.Println(jimmy.firstName)
    */

    bobby := person{
        firstName: "bobby",
        lastName: "tables",
        contact: contactInfo{
            email: "stuff@things.co",
            zipCode: 11111,
        },
    }
    fmt.Println(bobby.contact.email)
}
