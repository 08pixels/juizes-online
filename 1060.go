package main

import "fmt"

func main() {
	i := 0
	positives := 0
	number := 0.0

	for i < 6 {
		fmt.Scan(&number)
		if number > 0 {
			positives++
		}
		i++
	}

	fmt.Printf("%d valores positivos\n", positives)
}
