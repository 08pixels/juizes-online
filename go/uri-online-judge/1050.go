package main

import "fmt"

var city = map[int]string{
	61: "Brasilia",
	71: "Salvador",
	11: "Sao Paulo",
	21: "Rio de Janeiro",
	32: "Juiz de Fora",
	19: "Campinas",
	27: "Vitoria",
	31: "Belo Horizonte",
}

var ddd int

func main() {
	fmt.Scan(&ddd)
	result, ok := city[ddd]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("DDD nao cadastrado")
	}
}
