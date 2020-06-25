package main

import "fmt"

var dimension int
var matrix = [100][100]int{}

func isValid(u int, v int, limit int) bool {
	canX := u >= limit && u < (dimension-limit)
	canY := v >= limit && v < (dimension-limit)

	return canX && canY
}

func fillBorder(i int) {
	u, v := i, i

	for isValid(u, v, i) {
		matrix[u][v] = i + 1
		v++
	}
	v--
	for isValid(u, v, i) {
		matrix[u][v] = i + 1
		u++
	}
	u--
	for isValid(u, v, i) {
		matrix[u][v] = i + 1
		v--
	}
	v++
	for isValid(u, v, i) {
		matrix[u][v] = i + 1
		u--
	}

}

func main() {

	for true {
		fmt.Scan(&dimension)

		if dimension == 0 {
			break
		}

		for i := 0; i < dimension; i++ {
			fillBorder(i)
		}

		for i := 0; i < dimension; i++ {
			for j := 0; j < dimension; j++ {
				if j == 0 {
					fmt.Printf("%3d", matrix[i][j])
				} else {
					fmt.Printf("%4d", matrix[i][j])
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}
