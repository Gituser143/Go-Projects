package main

import "fmt"

func Swap(intArr []int, index int) {
	temp := intArr[index]
	intArr[index] = intArr[index+1]
	intArr[index+1] = temp
}

func BubbleSort(intArr []int) {
	for i := len(intArr); i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if intArr[j] > intArr[j+1] {
				Swap(intArr, j)
			}
		}
	}
}

func PrintSlice(intArr []int) {
	for i := 0; i < len(intArr); i++ {
		fmt.Printf("%d ", intArr[i])
	}
	fmt.Println()
}

func main() {
	numSlice := make([]int, 0)
	var num int
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanf("%d", &num)
		numSlice = append(numSlice, num)
	}

	BubbleSort(numSlice)

	PrintSlice(numSlice)
}
