package main

import "fmt"

var salary float64

func main() {
	fmt.Scan(&salary)

	if salary <= 2000 {
		fmt.Println("Isento")
	} else if salary <= 3000 {
		fmt.Printf("R$ %.2f\n", (salary-2000)*0.08)
	} else if salary <= 4500 {
		fmt.Printf("R$ %.2f\n", 80+(salary-3000)*0.18)
	} else {
		fmt.Printf("R$ %.2f\n", 80+(1500*0.18)+(salary-4500)*0.28)
	}
}
