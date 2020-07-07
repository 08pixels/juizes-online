package main

import "fmt"

var x, y int

func main() {

  for true {
    fmt.Scan(&x, &y)

    if x == 0 || y == 0 {
      break
    }

    if x > 0 && y > 0 {
      fmt.Println("primeiro")
    } else if x < 0 && y > 0 {
      fmt.Println("segundo")
    } else if x < 0 && y < 0 {
      fmt.Println("terceiro")
    } else {
      fmt.Println("quarto")
    }
  }

}
