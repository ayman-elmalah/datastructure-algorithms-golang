package main

import "fmt"

func main() {
	strings := []string{"Flowers", "Flows", "Flight"}

	fmt.Println(longestCommonPrefix(strings)) // Output: Fl
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	max := 0
	for {
		for i := 1; i < len(strs); i++ {
			if max >= len(strs[i-1]) || max >= len(strs[i]) || strs[i-1][max] != strs[i][max] {
				return strs[0][:max]
			}
		}
		max++
	}
}
