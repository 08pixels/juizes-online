package main

import "fmt"

var cases int
var player, choice [2]string
var value [2]int

func main() {
	fmt.Scan(&cases)

	for cases != 0 {
		for i := 0; i < 2; i++ {
			fmt.Scan(&player[i], &choice[i])
		}
		fmt.Scan(&value[0], &value[1])

		if choice[0] == "PAR" && (value[0]&1 == value[1]&1) {
			fmt.Println(player[0])
		} else if choice[0] == "IMPAR" && (value[0]&1 != value[1]&1) {
			fmt.Println(player[0])
		} else {
			fmt.Println(player[1])
		}

		cases--
	}
}
