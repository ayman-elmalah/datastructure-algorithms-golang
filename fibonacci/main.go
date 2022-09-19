package main

import "fmt"

func main() {
	// Find the number in the given index from fibonacci pattern    0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 ...
	index := 7
	fmt.Println(fibonacciRecursive(index)) // output : 13
	fmt.Println(fibonacciIterative(index)) // output : 13
}

func fibonacciRecursive(index int) int {
	if index <= 1 {
		return index
	}
	return fibonacciRecursive(index-1) + fibonacciRecursive(index-2)
}

func fibonacciIterative(index int) int {
	fibonacci := map[int]int{0: 0, 1: 1}

	if index > 1 {
		for i := 2; i <= index; i++ {
			fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
		}
	}

	return fibonacci[index]
}
