package main

import "fmt"

var inf, sup int

func main() {
  i := 0
  for i==0 {
    fmt.Scan(&inf, &sup)

    if inf == sup {
      break
    }

    if inf < sup {
      fmt.Println("Crescente")
    } else {
      fmt.Println("Decrescente")
    }
  }
}
