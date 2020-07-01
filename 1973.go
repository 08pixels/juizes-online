package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cases, i int
var total, number int64
var list []byte

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(list, 1000000*9)
	fmt.Scan(&cases)

	amount, steal := cases, cases
	min, flag := 0, true

	for scanner.Scan() {
		list := scanner.Text()

		numbers := strings.Fields(list)

		for _, v := range numbers {
			number, _ = strconv.ParseInt(v, 10, 64)
			total += number

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
	}

	fmt.Printf("%d %d\n", amount, total-int64(steal))
}
