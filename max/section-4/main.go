package main

import "fmt"

// structs are custom types
type person struct {
    firstName string
    lastName string

    // embedded structs!
    // makes sense, just another property/type pair
    // in this case, the propterty name and the struct name are not the same
    // contact contactInfo

    // this can also be done where the property name is the same as the struct
    contactInfo
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
        contactInfo: contactInfo{
            email: "stuff@things.co",
            zipCode: 11111,
        },
    }
    fmt.Println(bobby.contactInfo.email)
    bobby.print()

    // &variableName for access to pointers
    // pointer aka address in memory, NOT underlying value

    // the long way to make a pointer variable
    // bobbyPointer := &bobby


    // ALSO! go will let you call receivers for `*type` on a variable of `type`
    bobby.updateName("not bobby")

    bobby.print()
}

// structs can have receiver functions as well
func (p person) print() {
    fmt.Printf("%+v\n", p)
}

/*
// this does not actually update the existing firstName propterty
// "Go is a pass-by-value language"
// when calling a receiver, go makes a copy of the struct
// and performs any changes on the copy
// kind of like a mutable reference? 
func (p person) updateName(name string) {
    p.firstName = name
}
*/

// in order to update the underlying struct,
// pass a pointer to the struct
// `*pointer` is the value at the pointer, the data itself
// for function signatures, `*type` is a pointer to `type`
func (pointerToPerson *person) updateName (name string) {
    // take the pointer, change it to a value
    // and update the firstName property
    (*pointerToPerson).firstName = name
}


/*
001                 person{firstName: "idk"}
address/pointer     value
turn `address` into `value` using `*address`
turn `value` into `address` using `&value`
*/


/* reference types
types that include a reference/pointer to an underlying data structure
so there's no need to create one
slice - includes a reference to an underlying array
maps
channels
pointers
functions

value types
require pointers to change values
int
float
bool
string
struct
*/
