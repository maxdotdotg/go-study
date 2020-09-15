package main

import (
    "flag"
    "fmt"
    "io"
    "os"
)


func main(){
    /* 
    check for args!
    there's likely a library to handle this
    also, this was guesswork, I don't see anything in the docs
    I'm sure it's there, but I didn't find it
    maybe it's this?
    https://golang.org/pkg/os/#pkg-variables
    */

    // this time, let's use flag.Args()
    // https://golang.org/pkg/flag/#Args
    // flags.Args returns a slice of strings, so we can 
    // loop over it after it becomes available with flag.Parse()
    flag.Parse()

    /*
    open the file
    https://golang.org/pkg/os/#Open
    */
    for _, fileName := range flag.Args() {
        fileName, err := os.Open(fileName)
        if err != nil {
            fmt.Println("error", err)
            os.Exit(1)
        }
        io.Copy(os.Stdout, fileName)
    }

    /* 
    I still do not understand this...

    I THINK THIS WORKS BECAUSE
    we're passing a pointer to io.Copy... 
    os.Stdout implements Reader... and file implements writer?

    yep, os.Open implements Write, https://golang.org/pkg/os/#File.Write
    and io.Copy takes a pointer to something that implements Write

    "Stdin, Stdout, and Stderr are open Files pointing to the standard input,
    standard output, and standard error file descriptors."
    https://golang.org/pkg/os/#pkg-variables
    

    io.Copy(os.Stdout, file)
    */
}

