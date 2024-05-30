package main

import (
	"fmt"
)

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }

	/*
	for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }

	0
	10
	20
	30
	40
	50
	60
	70
	80
	90
	100
	*/

	/* range
	for index, value := array|slice|map {
   // code to be executed for each iteration
	}

	*/

	fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }

	fruits_1 := [3]string{"apple", "orange", "banana"}

  for idx, _ := range fruits_1 {
     fmt.Printf("%v\n", idx)
  }
	
}
