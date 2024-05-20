package main

import (
	"fmt"
)

// The Println() function is similar to Print() with the difference that a whitespace is
// added between the arguments, and a newline is added at the end:

// The Printf() function first formats its argument based on the given formatting verb and then prints them.


func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)

	//Output: HelloWorld	

	fmt.Print(i, "\n")
  fmt.Print(j, "\n")

	/* Output: HelloWorld	
		Hello
		World
	*/

	fmt.Print(i, "\n",j)

	/*
	Hello
	World
	*/


	var a string = "Hello"
  var b int = 15

  fmt.Printf("i has value: %v and type: %T\n", a, b)
  fmt.Printf("j has value: %v and type: %T", a, b)

	/*
	i has value: Hello and type: string
	j has value: 15 and type: int
	*/
}