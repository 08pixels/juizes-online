package main

import "fmt"

var number int

func main() {
  fmt.Scan(&number)

  if number > 2 {
    for i:=2; i<=10000; i+=number{
      fmt.Println(i)
    }
  }
}
