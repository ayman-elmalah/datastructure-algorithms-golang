package main

import "fmt"

func main() {
	items := []int{1, 4, 5, 8, 9, 17, 20, 23}

	fmt.Println(binarySearch(items, 8))  // Output: true
	fmt.Println(binarySearch(items, 54)) // Output: false
}

func binarySearch(items []int, needed int) bool {
	median := len(items) / 2

	switch {
	case len(items) == 0:
		return false
	case items[median] < needed:
		return binarySearch(items[median+1:], needed)
	case items[median] > needed:
		return binarySearch(items[:median], needed)
	default:
		return items[median] == needed
	}
}
