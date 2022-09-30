package main

import "fmt"

func main() {
	array := []int{8, 6, 1, 5, 4, 3}

	fmt.Println(bubbleSort(array)) // Output : [1 3 4 5 6 8]
}

func bubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
