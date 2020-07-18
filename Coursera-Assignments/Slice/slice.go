package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	intSlice := make([]int, 0, 3)
	var char string
	for {
		fmt.Println("Enter element to add to list (X to cancel)")
		fmt.Scanf("%s", &char)

		if strings.Compare(char, "X") == 0 {
			fmt.Println("\nSlice is", intSlice)
			break
		}

		i, _ := strconv.Atoi(char)
		intSlice = append(intSlice, i)

		sort.Ints(intSlice)
		fmt.Println("\nSlice is", intSlice)

	}
}
