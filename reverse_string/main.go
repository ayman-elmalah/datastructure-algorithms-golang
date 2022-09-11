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
	for i := 0; i <= (len(str)/2)-1; i++ {
		firstValue := string(str[i])
		secondValue := string(str[len(str)-1-i])
		str = replaceAtIndex(str, firstValue, len(str)-1-i)
		str = replaceAtIndex(str, secondValue, i)
	}

	return str
}

func replaceAtIndex(input string, replacement string, index int) string {
	return input[:index] + replacement + input[index+1:]
}
