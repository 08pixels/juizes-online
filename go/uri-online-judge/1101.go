package main

import "fmt"

var inf, sup int

func main(){
  j := 0
  for j == 0{
    fmt.Scan(&inf, &sup)

    if inf <= 0 || sup <= 0{
      break
    }
    if inf > sup{
      inf, sup = sup, inf
    }

    for i:=inf; i<=sup; i++{
      fmt.Printf("%d ", i)
    }
    fmt.Printf("Sum=%d\n", ((inf+sup)*(sup-inf+1))/2);
  }
}
