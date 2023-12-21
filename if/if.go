package main

import "fmt"

func main() {
  
  if 5%2 == 0 {
    fmt.Println("5 is even")
  } else {
    fmt.Println("5 is odd")
  }

  //A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
  if n := 9; n < 0 {
    fmt.Println(n, "is less than 0")
  } else if n < 10 {
    fmt.Println(n, "is single digit.")
  } else {
    fmt.Println(n, "is double digit")
  }
}
