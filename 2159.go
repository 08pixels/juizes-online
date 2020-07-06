package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scan(&number)

	fmt.Printf("%.1f %.1f\n", number/math.Log(number), 1.25506*(number/math.Log(number)))
}
