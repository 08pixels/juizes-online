package main

import "fmt"

var cases int
var numerator, denominator float64

func main() {
  fmt.Scan(&cases)
  for i:=0; i<cases; i++ {
    fmt.Scan(&numerator, &denominator)

    if denominator == 0.0{
      fmt.Println("divisao impossivel")
    } else {
      fmt.Printf("%.1f\n", numerator/denominator)
    }
  }

}
