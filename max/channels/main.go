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
       statusCheck(url)
    }
}

func statusCheck(url string) {
    // it's really odd to me that we aren't checking the return status
    // I don't like it
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(url, "might be down. err:", err)
        // add empty return statement to exit the function?
        return
    }
    fmt.Println(url, "is accepting traffic, got", resp.Status)
}
