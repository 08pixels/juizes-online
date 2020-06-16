package main

import "fmt"

var structure, group, foodChain string

func main() {
	fmt.Scan(&structure, &group, &foodChain)

	if structure == "vertebrado" {
		if group == "ave" {
			if foodChain == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if foodChain == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	} else {
		if group == "inseto" {
			if foodChain == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if foodChain == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
