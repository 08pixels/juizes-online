package main

import "fmt"

var hour, minutes int

func main() {

	for true {
		_, err := fmt.Scanf("%d:%d", &hour, &minutes)

		if err != nil {
			break
		}

		minutesTotal := hour*60 + minutes
		target := 7 * 60

		if minutesTotal <= 7*60 {
			fmt.Printf("Atraso maximo: %d\n", 0)
		} else {
			fmt.Printf("Atraso maximo: %d\n", minutesTotal-target)
		}
	}
}
