package main

import "fmt"

var cases int
var word1, word2 string

func max(size1 int, size2 int) int {

	if size1 > size2 {
		return size1
	}

	return size2
}

func main() {

	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&word1, &word2)
		answer := []byte{}

		size1 := len(word1)
		size2 := len(word2)

		for i := 0; i < max(size1, size2); i++ {
			if i < size1 {
				answer = append(answer, word1[i])
			}
			if i < size2 {
				answer = append(answer, word2[i])
			}
		}

		fmt.Println(string(answer))
		cases--
	}
}
