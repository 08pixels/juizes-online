package main

import "fmt"

func main() {
	j := 0
	i := -1

	for i < 2 {
		j %= 10

		if j == 0 {
			i++
			fmt.Printf("I=%d J=%d\n", i, 1+i)
			fmt.Printf("I=%d J=%d\n", i, 2+i)
			fmt.Printf("I=%d J=%d\n", i, 3+i)
		} else {
			fmt.Printf("I=%d.%d J=%d.%d\n", i, j, 1+i, j)
			fmt.Printf("I=%d.%d J=%d.%d\n", i, j, 2+i, j)
			fmt.Printf("I=%d.%d J=%d.%d\n", i, j, 3+i, j)
		}

		j += 2
	}
}
