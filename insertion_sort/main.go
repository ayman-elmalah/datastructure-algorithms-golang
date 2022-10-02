package main

import "fmt"

func main() {
	array := []int{5, 4, 7, 9, 1, 8, 3, 2}

	fmt.Println(insertionSort(array)) // Output: [1 2 3 4 5 7 8 9]
}

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}

	return array
}
