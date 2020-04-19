package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}

	for _, url := range websites {
		err := checkUrlStatus(url)

		if err != nil {
			fmt.Println("The Website", url, "is down")
		} else {
			fmt.Println(url, "is up and running")
		}
	}
}

func checkUrlStatus(url string) error {
	_, err := http.Get(url)
	return err
}
