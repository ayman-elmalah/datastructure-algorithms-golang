package main

import "fmt"

func main() {
	array := []int{5, 2, 0, 8, 3}

	fmt.Println(selectionSort(array))
}

func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		smallestIndex := i
		for j := i; j < len(array); j++ {
			if array[j] < array[smallestIndex] {
				smallestIndex = j
			}
		}
		array[i], array[smallestIndex] = array[smallestIndex], array[i]
	}

	return array
}
