## general
`go run` compiles and executes the code in the working dir
`go build` compiles the code in the working dir, not running anything
`go fmt` linting, including 8 spaces for indentation
`go install` compile and install (make availabile via path?)
`go get` get a library or similar from the internet
`go test` run tests, yo

when making an executable, instead of a lib, the package is always named main, and _must_ have a function called main


## slices
- arrays and slices must contain data of one type
- slices are mutable
- arrays are immutable
- writing to a file requires writing bytes, not strings, so srings must be cast/coverted/something to bytes
    yep, using type conversion: `targetType(currentType)` so `[]byte("hello human")`


## maps
- types of keys can't change
- types of values can't change either
- differences between structs and maps
    structs can have different types of keys and values
    map keys are indexed, struct keys are not
    maps are reference types, structs are value types
    can't add new properties to structs on the fly, they must all be known at compile time


## interfaces
- one goal is to make code reuse easier?
- function overloading is not a thing in go
- interfaces aren't concrete types like int or struct
    they're an implicit superset of concrete types? I think?

