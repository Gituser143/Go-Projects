package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		log.Fatal(err)
	}

	body := make([]byte, 100000)

	n, err2 := resp.Body.Read(body)
	if n == 0 && err2 != nil {
		log.Fatal(err2)
	}

	str := string(body)
	fmt.Println(str)
}
