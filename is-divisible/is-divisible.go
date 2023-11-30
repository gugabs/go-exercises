package main

import "fmt"

func main() {
  divider := 0

  fmt.Println("Choose the divider:")
  fmt.Scanf("%d", &divider)

  for i := 1; i <= 100; i++ {
    if isDivisible(i, divider) {
      fmt.Printf("O número %d é divisível por %d\n", i, divider)
    }
  }
}

func isDivisible(n int, divider int) bool {
  return n%divider == 0
}
