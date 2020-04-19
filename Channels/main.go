package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://twitter.com"}

	c := make(chan string)

	for _, url := range websites {
		go checkUrlStatus(url, c)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}
}

func checkUrlStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		// fmt.Println("The Website", url, "is down")
		c <- "The Website " + url + " is down"
	} else {
		// fmt.Println(url, "is up and running")
		c <- url + " is up and running"
	}
}
