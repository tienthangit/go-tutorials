/* Syntax
slice_name := []datatype{values}

A common way of declaring a slice is like this:
myslice := []int{}

VD: myslice := []int{1,2,3}

len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
*/

package main

import (
	"fmt"
)

func _main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)

	/*
	0
	0
	[]
	4
	4
	[Go Slices Are Powerful]
	*/
}