package main

import "fmt"

func main() {
	var tabs, activity int
	var action string

	fmt.Scan(&tabs, &activity)

	for activity != 0 {
		fmt.Scan(&action)

		if action == "fechou" {
			tabs++
		} else {
			tabs--
		}
		activity--
	}
	fmt.Println(tabs)
}
