package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Naive approach
	// for i := 0; i < len(websites); i++ {
	// 	fmt.Println(<-c)
	// }

	// Better than Naive approach
	// for {
	// 	go checkUrlStatus(<-c, c)
	// }

	// Cleaner approach
	// for url := range c {
	// 	go checkUrlStatus(url, c)
	// }

	// Safer approach (Doesn't DoS the sites)
	for url := range c {
		go func(url2 string) {
			time.Sleep(time.Second * 5)
			checkUrlStatus(url2, c)
		}(url)
	}
}

func checkUrlStatus(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("The Website", url, "is down")
		fmt.Println(err)
		c <- url
	} else {
		fmt.Println(url, "is up and running")
		c <- url
	}
}
