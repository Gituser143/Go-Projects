package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	websites := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}

	for _, url := range websites {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("Error: Could not fetch", url)
		} else {
			fmt.Println("\n======================\n", url, "\n======================\n ")
			io.Copy(os.Stdout, resp.Body)
		}
	}
}
