package main

import "fmt"

func add(x, y int) int {

  sum := x + y
  return sum
}

func main() {
  
  fmt.Println("sum:", add(2, 3))
}
