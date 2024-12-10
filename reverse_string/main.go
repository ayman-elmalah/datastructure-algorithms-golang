package main

import "fmt"

func main() {
	fmt.Println(reverseStringLoop("Hello world"))     // Using loop    			// Output: dlrow olleH
	fmt.Println(reverseStringHalfLoop("Hello world")) // Using half loop    	// Output: dlrow olleH
}

// O(n)
func reverseStringLoop(str string) string {
	var newString = ""

	for i := len(str) - 1; i >= 0; i-- {
		newString = newString + string(str[i])
	}

	return newString
}

// O(n)
func reverseStringHalfLoop(str string) string {
	runes := []rune(str)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}

	return string(runes)
}