package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		log.Fatal(err)
	}

	// body := make([]byte, 100000)
	//
	// n, err2 := resp.Body.Read(body)
	// if n == 0 && err2 != nil {
	// 	log.Fatal(err2)
	// }
	//
	// str := string(body)
	// fmt.Println(str)

	// The below line is a simpler substitute for the above comments
	// io.Copy(os.Stdout, resp.Body)

	// THe below few lines uses a custom writer as destination
	wirterVar := customWriter{}
	io.Copy(wirterVar, resp.Body)
}

func (customWriter) Write(p []byte) (int, error) {
	_, err := fmt.Println(string(p))
	return len(p), err
}
