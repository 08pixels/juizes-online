package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cases, i int
var total int64

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)

	amount := cases
	steal := cases
	min := 0
	flag := true

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		fmt.Println(number)

		total = int64(number)

		if number == 1 {
			min = i + 1
		}

		if flag && number&1 == 0 {
			flag = false
			amount = i + 1
			steal = amount + (i - min)
		}
		i++
	}

	fmt.Printf("%d %d\n", amount, total-int64(steal))
}
