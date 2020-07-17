package main

import "fmt"

func main() {
	var f float64
	for x := 0; x < 2; x++ {
		fmt.Println("Enter float number")
		fmt.Scanf("%f", &f)
		i := int(f)
		fmt.Printf("Truncated integer is %d\n\n", i)
	}
}
