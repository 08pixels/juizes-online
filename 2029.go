package main

import (
	"fmt"
	"math"
)

func main() {

	var volume, diameter float64
	for true {
		_, err := fmt.Scan(&volume, &diameter)

		if err != nil {
			break
		}
		circunrefenceArea := 3.14 * math.Pow(diameter/2, 2)

		fmt.Printf("ALTURA = %.2f\n", volume/circunrefenceArea)
		fmt.Printf("AREA = %.2f\n", circunrefenceArea)
	}

}
