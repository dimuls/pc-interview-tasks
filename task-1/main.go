package main

import "fmt"

func main() {
	array := []int{1, 3, 1, 3, 1, 3, 8, 8, 3}
	k := 1

	fmt.Println(doubles(array, k))
}

func doubles(array []int, k int) map[int]int {
	doubles := map[int]int{}
	jMax := len(array)
	for i := range array {
		j := i + k + 1
		if j >= jMax {
			break
		}
		if array[i] == array[j] {
			doubles[array[i]]++
		}
	}
	return doubles
}
