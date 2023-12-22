package main

import (
  "fmt"
  "maps"
)

func main() {
  
  // make is used to create an empty object with non-zero length
  m := make(map[string]int)

  m["k1"] = 2
  m["k2"] = 4

  fmt.Println("map:", m)

  // Zero value for the value type is returned
  // the value type is specified on line 11 (pay attention)
  v1 := m["k1"]
  fmt.Println("v1:", v1)

  // delete k:v pair
  delete(m, "k1")

  // remove all key value pairs
  clear(m)
  fmt.Println("m:", m)

  // check if key exists or just lucky Zero Value type matches key type (check line 18)
  _, prs := m["k1"]
  fmt.Println("prs:", prs)


  n := map[string]int{"foo": 1, "bar": 2}
  m2 := map[string]int{"foo": 1, "bar": 2}

  if maps.Equal(n, m2) {
    fmt.Println("n == m2")
  }
}
