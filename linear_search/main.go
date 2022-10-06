package main

import "fmt"

func main() {
	items := []int{5, 4, 7, 9, 1, 8, 3, 2}

	fmt.Println(linearSearch(items, 8))  // Output: true
	fmt.Println(linearSearch(items, 54)) // Output: false
}

func linearSearch(items []int, key int) bool {
	for _, item := range items {
		if item == key {
			return true
		}
	}

	return false
}
