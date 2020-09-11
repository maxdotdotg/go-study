package main

import (
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
    args := os.Args
    if len(args) != 2 {
        fmt.Println("accepts only one argument, a file to print to stdout")
        os.Exit(1)
    }

    /*
    open the file
    https://golang.org/pkg/os/#Open
    */
    file, err := os.Open(string(args[1]))
    if err != nil {
        fmt.Println("error", err)
        os.Exit(1)
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
    

    */
    io.Copy(os.Stdout, file)
}

