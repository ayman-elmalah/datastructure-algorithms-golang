package main

import "fmt"

func main() {
	fmt.Println(findMissingNUmber([]int{2, 4, 6, 3, 7, 8}))  // Expected: 5
	fmt.Println(findMissingNUmber([]int{2, 4, 5, 6}))        // Expected: 3
	fmt.Println(findMissingNUmber([]int{2, 3, 4, 5, 7}))     // Expected: 6
	fmt.Println(findMissingNUmber([]int{1, 3, 5, 4}))        // Expected: 2
	fmt.Println(findMissingNUmber([]int{-5, -4, -3, -1, 0})) // Expected: -2
}

func findMissingNUmber(n []int) int {
	numbers := sort(n)
	selectedNumber := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] == selectedNumber+1 {
			selectedNumber++
		} else {
			return numbers[i] - 1
		}
	}

	return selectedNumber
}

func sort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}
