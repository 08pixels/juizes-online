package main

import "fmt"

var number int
var value = []int{1, 5, 10, 50, 100, 500, 1000}
var symbol = []string{"I", "V", "X", "L", "C", "D", "M"}

func main() {
	fmt.Scan(&number)
	result := ""
	positional := 1

	for number != 0 {
		val := (number % 10) * positional

		for i := 0; i < 7; i++ {
			if value[i] == val {
				result = symbol[i] + result
				val = 0
			}

			for j := i + 1; j < 7; j++ {
				if value[j]-value[i] == val {
					result = symbol[i] + symbol[j] + result
					val = 0
				}
			}
		}

		if val != 0 {
			complement := ""
			for i := 6; i >= 0; i-- {
				for value[i] <= val {
					complement += symbol[i]
					val -= value[i]
				}
			}
			result = complement + result
		}

		number /= 10
		positional *= 10
	}

	fmt.Println(result)
}
