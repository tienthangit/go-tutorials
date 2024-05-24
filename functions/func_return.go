/*

func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}

*/

package main

import (
	"fmt"
)

func myFunction(x int, y int) int {
  return x + y
}

func myFunction1(x int, y int) (result int) {
  result = x + y
  return
}

func myFunction2(x int, y int) (result int) {
  result = x + y
  return result
}

func myFunction3(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func myFunction4(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func myFunction5(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}


func main() {
  fmt.Println(myFunction3(5, "Hello")) // Hello world!

  a, b := myFunction4(5, "Hello")
  fmt.Println(myFunction4(5, "Hello"))

  _, b := myFunction5(5, "Hello")
  fmt.Println(b) // Hello
}
