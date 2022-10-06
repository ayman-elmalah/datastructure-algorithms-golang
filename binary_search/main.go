package main

import "fmt"

func main() {
	items := []int{1, 4, 5, 8, 9, 17, 20, 23}

	fmt.Println(binarySearch(items, 8))  // Output: true
	fmt.Println(binarySearch(items, 54)) // Output: false
}

func binarySearch(items []int, key int) bool {
	median := len(items) / 2
	switch {
	case len(items) == 0:
		return false
	case items[median] == key:
		return true
	case items[median] > key:
		return binarySearch(items[:median], key)
	case items[median] < key:
		return binarySearch(items[median+1:], key)
	default:
		return false
	}
}
