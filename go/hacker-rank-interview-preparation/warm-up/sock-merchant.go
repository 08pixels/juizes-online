package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
    socks := [101]int32{}

    for i:=int32(0); i<n; i++ {
        socks[ar[i]]++
    }

    answer := int32(0)
    for i := 1; i <= 100; i++ {
        answer += (socks[i] / 2)
    }

    return answer

}

func main() {
  var number, x int32
  var ar []int32
  
  fmt.Scan(&number)
  
  for i:=0; i<number; i++ {
    fmt.Scan(&x)
    ar = append(ar, x)
  }
}
