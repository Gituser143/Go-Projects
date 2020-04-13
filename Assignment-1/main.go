package main

import "fmt"

func main() {
	arr := []int{}

	for i := 0; i <= 10; i++ {
		arr = append(arr, i)
	}

	for _, i := range arr {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
