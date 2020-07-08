package main

import (
	"fmt"
	"sort"
)

func copy(slice []int) []int {
	return slice
}

func main() {
	slice := make([]int, 3, 3)
	before := make([]int, 3, 3)

	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])
		before[i] = slice[i]
	}

	sort.Ints(slice)
	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println("")

	for _, v := range before {
		fmt.Println(v)
	}
}
