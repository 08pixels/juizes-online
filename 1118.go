package main

import "fmt"

func main() {
	option := 0

	for option != 2 {
		option = 0
		counter := 0
		testGrade, sum := 0.0, 0.0

		for counter != 2 {
			fmt.Scan(&testGrade)

			if testGrade < 0 || testGrade > 10 {
				fmt.Println("nota invalida")
			} else {
				sum += testGrade
				counter++
			}
		}

		fmt.Printf("media = %.2f\n", sum/2)

		for option != 1 && option != 2 {
			fmt.Println("novo calculo (1-sim 2-nao)")
			fmt.Scan(&option)
		}
	}

}
