package main

import "fmt"

func main() {
	counter, sum, testGrade := 0, 0.0, 0.0

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

}
