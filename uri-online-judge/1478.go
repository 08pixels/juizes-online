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
	u, v := i, i-1

	for isValid(u, v+1, i) {
		v++
		matrix[u][v] = v - i + 1
	}
	for isValid(u+1, v, i) {
		u++
		matrix[u][v] = matrix[u-1][v] - 1
	}
	for isValid(u, v-1, i) {
		v--
		matrix[u][v] = matrix[u][v+1] + 1
	}
	for isValid(u-1, v, i) {
		u--
		matrix[u][v] = matrix[u+1][v] - 1
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
