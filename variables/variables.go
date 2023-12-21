package main

import "fmt"

func main() {
  var a = "Jello"
  fmt.Println(a)

  // unlike python, types are specified after variable name
  // Go will infer type of initialized variables (initializing is setting a value, declaring is just.. declaring).
  var b, c int = 1, 2
  fmt.Println(b,c)

  // := shorthand for declaring + initializing
  f := "apple"
  fmt.Println(f)
}
