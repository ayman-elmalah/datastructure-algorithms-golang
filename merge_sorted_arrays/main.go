package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeSortedArrays([]int{0, 3, 6, 7}, []int{1, 8, 9})) // Output : [0 1 3 6 7 8 9]
}

func mergeSortedArrays(array1 []int, array2 []int) []int {
	var newArray []int

	for {
		if len(array1) == 0 {
			newArray = append(newArray, array2...)
		}

		if len(array2) == 0 {
			newArray = append(newArray, array1...)
		}

		if len(array1) == 0 || len(array2) == 0 {
			break
		}

		first := array1[0]
		second := array2[0]

		if first < second {
			newArray = append(newArray, first)
			array1 = RemoveFromArrayByIndex(array1, 0)
		} else {
			newArray = append(newArray, second)
			array2 = RemoveFromArrayByIndex(array2, 0)
		}
	}

	return newArray
}

func RemoveFromArrayByIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
