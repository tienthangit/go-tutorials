package main

import (
	"fmt"
)

// Typed Constants
const A int = 3

// UnTyped Constants
const B = 1

// Multiple Constants Declaration

const (
  C int = 1
  D = 3.14
  E = "Hi!"
)



func main() {
  fmt.Println(A)
	fmt.Println(B)

	fmt.Println(C)
  fmt.Println(D)
  fmt.Println(E)
}