package main

import (
    "fmt"
)

func main() {
    // this looks like a slice, so that's... a thing
    colors := map[string]string{
        "red": "#ff0000",
        "green": "#008000",
        // trailing commas, yo
    }
    fmt.Println(colors)

    // regular variable syntax
    // an empty map with empty zero values
    var coloring_books map[string]string
    fmt.Println(coloring_books)

    /* 
    // this won't work
    coloring_books["batman vs dracula"] = "a lot of black and red"
    // "it doesn't point to an initialized map. A nil map behaves like an empty
    // map when reading, but attempts to write to a nil map will cause a runtime
    // panic; don't do that. To initialize a map, use the built in make function"
    // https://blog.golang.org/maps
    */

    // another way to create empty maps
    cool_books := make(map[string]string)
    // add a key/value pair
    cool_books["poison ivy and harley quinn"] = "the best couple"
    cool_books["batman vs dracula"] = "a lot of black and red"
    fmt.Println(cool_books)

    // I'm wrong, batman vs dracula isn't actually cool
    delete(cool_books, "batman vs dracula")
    fmt.Println(cool_books)

    // also for fun, we can delete keys that don't exist without error
    delete(colors, "batman vs dracula")

}
