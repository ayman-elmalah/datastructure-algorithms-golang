package main

import (
	"errors"
	"fmt"
)

func main() {
	array := []int{1, 2, 2, 8, 12, 9, 5}

	fmt.Println(whichRepeatedFirstWithLooping(array))   // Using loops        // Output: 2 <nil>
	fmt.Println(whichRepeatedFirstWithHashTable(array)) // Using hashtable    // Output: 2 <nil>
}

// With looping => O(n^2)
func whichRepeatedFirstWithLooping(array []int) (int, error) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				return array[i], nil
			}
		}
	}

	return 0, errors.New("Undefined")
}

// With Hashtable => O(n)
func whichRepeatedFirstWithHashTable(array []int) (int, error) {
	newMap := make(map[int]int)

	for i := 0; i < len(array); i++ {
		_, ok := newMap[array[i]]

		if ok {
			return array[i], nil
		} else {
			newMap[array[i]] = array[i]
		}
	}

	return 0, errors.New("Undefined")
}
