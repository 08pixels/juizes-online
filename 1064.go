package main

import "fmt"

var number float64
var sum float64
var positives int

func main() {
	i := 0

	for i < 6 {
		fmt.Scan(&number)

		if number > 0 {
			sum += number
			positives++
		}
		i++
	}
	fmt.Printf("%d valores positivos\n%.1f\n", positives, sum/float64(positives))
}
