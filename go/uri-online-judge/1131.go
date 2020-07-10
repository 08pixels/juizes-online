package main

import "fmt"

var goalInter, goalGremio int
var winInter, winGremio, draw int

func main() {
	option := 0

	for option != 2 {
		fmt.Scan(&goalInter, &goalGremio)

		if goalInter > goalGremio {
			winInter++
		} else if goalInter < goalGremio {
			winGremio++
		} else {
			draw++
		}

		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scan(&option)
	}

	fmt.Printf("%d grenais\n", winInter+winGremio+draw)
	fmt.Printf("Inter:%d\n", winInter)
	fmt.Printf("Gremio:%d\n", winGremio)
	fmt.Printf("Empates:%d\n", draw)

	if winInter > winGremio {
		fmt.Println("Inter venceu mais")
	} else if winInter < winGremio {
		fmt.Println("Gremio venceu mais")
	} else {
		fmt.Println("Nao houve vencedor")
	}
}
