package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

type logWriter struct{}

func main() {
    resp, err := http.Get("http://google.com")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    /* 

    ALL OF THIS IS THE VERY LONG WAY TO DO IT

    // make a slice of bytes that has 99_999 empty values
    // `Read` doesn't re-size the slice, so let's make it big
    // to avoid losing that may not fit in the slice of bytes
    bs := make([]byte, 99999)

    // something about reference types and value types allows
    // this to work... `Read` updates the underlying slice
    // with the content of Body... because of this, we don't
    // need to store the value in another variable
    resp.Body.Read(bs)
    // we do have to cast it to a string, because otherwise 
    // it's just bytes and that's crap for humans to read
    fmt.Println(string(bs))

    */

    /* potentially the short way?
    // it is organized as dest first, source second, that's annoying
    // copy `resp.Body` and write it to stdout
    // os.Stdout implements the Writer interface
    // resp.Body implements the Reader interface
    io.Copy(os.Stdout, resp.Body)
    */

    lw := logWriter{}
    io.Copy(lw, resp.Body)
}


func (logWriter) Write(bs []byte) (int, error) {
    /*
    // just because you implement the interface doesn't
    // mean you did it right, so PAY ATTENTION
    // this technically implements the Writer interface
    // but it doesn't do what we actually want
    return 1, nil
    */

    fmt.Println(string(bs))
    fmt.Println("we wrote some bytes! this many:", len(bs))
    return len(bs), nil
}
