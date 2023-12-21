package main

import (
  "fmt"
  "math"
)

const s string =  "constant"

func main() {
  fmt.Println(s)

  // Doesn't have a type until explicitly cast
  // or used in a particular context
  const n = 5000000
  const d = 3e20 / n
  fmt.Println(d)

  // Only Uppercased names are exported in a package
  // Lower case is hidden
  fmt.Println(math.Sin(n))
}
