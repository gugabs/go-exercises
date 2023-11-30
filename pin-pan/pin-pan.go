package main

import "fmt"

func main() {
  fmt.Println("Pin-Pan!")

  for i := 0; i <= 100; i++ {
    shouldPrintCurrN := true

    if isMultipleOfThree(i) {
      fmt.Println("Pin")
      shouldPrintCurrN = false
    }

    if isMultipleOfFive(i) {
      fmt.Println("Pan")
      shouldPrintCurrN = false
    }

    if shouldPrintCurrN {
      fmt.Println(i)
    }
  }
}

func isMultipleOfThree(n int) bool {
  return n%3 == 0
}

func isMultipleOfFive(n int) bool {
  return n%5 == 0
}
