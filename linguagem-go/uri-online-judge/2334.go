package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number string

	for true {
		fmt.Scan(&number)

		if number == "-1" {
			break
		}

		littleDucks, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			break
		}

		if littleDucks == uint64(0) {
			littleDucks = uint64(1)
		}

		fmt.Println(littleDucks - uint64(1))
	}
}
