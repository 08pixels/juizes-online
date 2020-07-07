package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	var answer int32
	current := 0

	for current != len(c)-1 {
		if c[current+2] == 0 {
			current += 2
		} else if c[current+1] == 0 {
			current++
		}
		answer++
	}

	return answer
}

func main() {
	var number, x int32
	var clouds []int32

	fmt.Scan(&number)

	for i := 0; i < int(number); i++ {
		fmt.Scan(&x)
		clouds = append(clouds, x)
	}

	fmt.Println(jumpingOnClouds(clouds))
}
