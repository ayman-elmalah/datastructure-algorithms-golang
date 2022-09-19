package main

import "fmt"

func main() {
	number := 6
	fmt.Println(factorialRecursive(number)) // output : 720  // O(n)
	fmt.Println(factorialIterative(number)) // output : 720  // O(n)
}

func factorialRecursive(number int) int {
	if number == 1 || number == 0 {
		return number
	}
	return number * factorialRecursive(number-1)
}

func factorialIterative(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}
