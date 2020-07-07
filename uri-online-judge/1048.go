package main

import "fmt"

var salary float64

type ratio struct {
	high       float64
	porcentage int
}

var table = []ratio{
	ratio{
		high: 400, porcentage: 15,
	},
	ratio{
		high: 800, porcentage: 12,
	},
	ratio{
		high: 1200, porcentage: 10,
	},
	ratio{
		high: 2000, porcentage: 7,
	},
	ratio{
		high: 1000000000000, porcentage: 4,
	},
}

func main() {
	fmt.Scan(&salary)
	for _, data := range table {

		if salary <= data.high {
			fmt.Printf("Novo salario: %.2f\n", salary*(1+float64(data.porcentage)/100))
			fmt.Printf("Reajuste ganho: %.2f\n", salary*(float64(data.porcentage)/100))
			fmt.Printf("Em percentual: %d %%\n", data.porcentage)

			break
		}
	}
}
