package main

import (
    "fmt"
    "net/http"
)

func main() {
    links := []string {
        "http://www.google.com",
        "http://www.amazon.com",
        "http://www.boingboing.net",
    }


    for _, url := range links {
        // when prefixed with `go`, start this line of code
        // in a separate go routine
        // when a go routine enters a blocking state, go will
        // run the subsequent go routine
        // tl;dr, run until blocking, then start the next one
        go statusCheck(url)
    }

    /*
    // serial version
    for _, url := range links {
       statusCheck(url)
    }
    */
}

func statusCheck(url string) {
    // this is a blocking call, the go routine that executes main can't do
    // anything until this is completed
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(url, "might be down. err:", err)
        // add empty return statement to exit the function?
        return
    }
    fmt.Println(url, "is accepting traffic, got", resp.Status)
}

