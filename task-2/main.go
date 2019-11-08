package main

import "fmt"

func main() {
	array := []int{1, 3, 1, 3, 1, 3, 8, 8, 3, 2}

	fmt.Println(haveSum(array, 11))
}

func haveSum(array []int, sum int) bool {
	m := map[int]struct{}{}
	for _, v := range array {
		if _, exists := m[sum-v]; exists {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}
