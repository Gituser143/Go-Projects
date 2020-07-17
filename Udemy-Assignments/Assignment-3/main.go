package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	files := os.Args[1:]

	for _, file := range files {
		fileToPrint, err := os.Open(file)

		if err != nil {
			fmt.Println("Error:", err)
		}

		io.Copy(os.Stdout, fileToPrint)
	}
}
