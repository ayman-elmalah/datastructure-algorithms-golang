package main

import (
	"fmt"
)

func main() {
	fmt.Println(Merge2SortedArrays([]int{0, 3, 6, 7}, []int{1, 8, 9})) // Output : [0 1 3 6 7 8 9]
}

func Merge2SortedArrays(arr1 []int, arr2 []int) []int {
	var result []int

	for {
		if len(arr1) == 0 {
			result = append(result, arr2...)
		}

		if len(arr2) == 0 {
			result = append(result, arr1...)
		}

		if len(arr1) == 0 || len(arr2) == 0 {
			break
		}

		if arr1[0] < arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}

	return result
}
