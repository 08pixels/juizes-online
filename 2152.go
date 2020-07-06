package main

import "fmt"

func main() {
	var cases int
	var hour, minute, action int
	fmt.Scan(&cases)

	for cases != 0 {
		fmt.Scan(&hour, &minute, &action)

		state := "abriu"
		if action == 0 {
			state = "fechou"
		}
		fmt.Printf("%02d:%02d - A porta %s!\n", hour, minute, state)
		cases--
	}
}
