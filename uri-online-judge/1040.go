package main

import "fmt"

var a, b, c, d float32

func main() {
	fmt.Scan(&a, &b, &c, &d)
	ponderedMean := (a*2 + b*3 + c*4 + d) / 10

	fmt.Printf("Media: %.1f\n", ponderedMean)

	if ponderedMean >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if ponderedMean < 5.0 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")
		fmt.Scan(&a)
		ponderedMean = (ponderedMean + a) / 2.0

		fmt.Printf("Nota do exame: %.1f\n", a)
		if ponderedMean >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", ponderedMean)
	}
}
