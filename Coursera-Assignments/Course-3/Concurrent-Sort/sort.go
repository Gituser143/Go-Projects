package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortArr(arr []int, wg *sync.WaitGroup) {
	sort.Ints(arr)
	fmt.Println("Sorted portion is", arr)
	wg.Done()
}

func merge(left, right []int, wg *sync.WaitGroup) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	wg.Done()
	return
}

func main() {
	intSlice := []int{}
	size := 0
	for {
		var input string
		fmt.Println("Enter element. X to finish")
		fmt.Scanf("%s", &input)
		if strings.Compare(input, "X") == 0 {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Breaking")
			break
		}

		intSlice = append(intSlice, num)
		size++
	}

	var wg sync.WaitGroup

	wg.Add(4)
	go sortArr(intSlice[:len(intSlice)/4], &wg)
	go sortArr(intSlice[len(intSlice)/4:len(intSlice)/2], &wg)
	go sortArr(intSlice[len(intSlice)/2:3*len(intSlice)/4], &wg)
	go sortArr(intSlice[3*len(intSlice)/4:], &wg)
	wg.Wait()

	wg.Add(2)
	first := merge(intSlice[:size/4], intSlice[size/4:size/2], &wg)
	second := merge(intSlice[size/2:3*size/4], intSlice[3*size/4:], &wg)
	wg.Wait()

	wg.Add(1)
	sortedArr := merge(first, second, &wg)
	wg.Wait()

	fmt.Println("Sorted Array is", sortedArr)
}
