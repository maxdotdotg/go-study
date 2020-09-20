package main

import (
	"fmt"
	"net/http"
    "time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.boingboing.net",
	}

	// make a channel to pass msgs between go routines
	// clearly, keeping track of all the short named variables is _easy_
	// pshhh plebs

	// make a channel of type string
	c := make(chan string)

	for _, url := range links {
		// when prefixed with `go`, start this line of code
		// in a separate go routine
		// when a go routine enters a blocking state, go will
		// run the subsequent go routine
		// tl;dr, run until blocking, then start the next one,
		// only one executes at once
		go statusCheck(url, c)
	}

	// waiting to receive data on a channel is also a blocking task
	// the program will exit when it receives data, even though there's
	// more to do, since it's waiting for a message, not all messages
	//fmt.Println(<- c)

	// this kind of upsets me, and I need to think about why
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

    // make that loop infinite!
    // for {
    //  DO SOME STUFF
    // }

    /* 
    // make the infinite loop more readable
    // for messages in channel (infinite because there's no limit to the number
    // of messages?)
    for l := range c {
        // send the content of channel c as input to the statusCheck 
        // function and make it a go routine, since statusCheck uses 
        // http.Get whicih is a blocking operation
        go statusCheck(l, c)
    }
    */

	/*
	   // serial version
	   for _, url := range links {
	      statusCheck(url)
	   }
	*/

    // use a function literal (lambda function from python) to insert a pause
    // between go routines
    // don't use variables declared outisde the function literal inside the
    // function literal, except the channel in this case, I guess?
    // not using variables from outside the function literal makes sure the
    // child go routines are accessing the correct data in memory, 
    for l := range c {
        go func(link string) {
            time.Sleep(5 * time.Second)
            statusCheck(link, c)
        }(l) // requires `()` at the end so it gets called
    }
}

func statusCheck(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down. err:", err)
		//c <- "is down maybe"
        // send the url to the channel instead
        c <- url
		return
	}
	fmt.Println(url, "is accepting traffic, got", resp.Status)
	// c <- "is up"
    c <- url
}

/*
non-channel version
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
*/
