package main

import "fmt"

func main() {
	array := []int{5, 4, 7, 9, 1, 8, 3, 2}

	fmt.Println(mergeSort(array)) // Output: [1 2 3 4 5 7 8 9]
}

// Big O notation O(n log(n))
func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	left := mergeSort(array[:len(array)/2])
	right := mergeSort(array[len(array)/2:])

	return merge(left, right)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
