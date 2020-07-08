
package main

import "fmt"

const password = 2002
var tempt int

func main() {
  denied := true
  for denied {
    fmt.Scan(&tempt)

    if tempt == password {
      fmt.Println("Acesso Permitido")
      denied = false
    } else {
      fmt.Println("Senha Invalida")
    }
  }
}

