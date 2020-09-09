package main

import (
    "fmt"
    "os"
)

type logWriter struct{}

func main(){
    // open the file
    file, err := os.Open(string(os.Args))
    if err != nil {
        fmt.Println("error", err)
        os.Exit(1)
    }

    // print the file
    fmt.Println(*file)

    // can't use this yet, since file returns a pointer
    //io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    return len(bs), nil
}
