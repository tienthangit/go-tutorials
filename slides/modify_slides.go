package main

import (
	"fmt"
)

/*
	prices use access elements to arrays
	You can append elements to the end of a slice using the append()function:

*/

func main() {
  prices := []int{10,20,30}

  prices[2] = 50
  fmt.Println(prices[0])
  fmt.Println(prices[2])

	// myslice1 := []int{1, 2, 3, 4, 5, 6}
  // fmt.Printf("myslice1 = %v\n", myslice1)
  // fmt.Printf("length = %d\n", len(myslice1))
  // fmt.Printf("capacity = %d\n", cap(myslice1))

  // myslice1 = append(myslice1, 20, 21)
  // fmt.Printf("myslice1 = %v\n", myslice1)
  // fmt.Printf("length = %d\n", len(myslice1))
  // fmt.Printf("capacity = %d\n", cap(myslice1))

	/*
		myslice1 = [1 2 3 4 5 6]
		length = 6
		capacity = 6
		myslice1 = [1 2 3 4 5 6 20 21]
		length = 8
		capacity = 12
	*/


	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice2 := arr1[1:3] // Slice array
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))

	/* if the array is large and you need only a few elements, 
	it is better to copy those elements using the copy() function.

	 The copy() function creates a new underlying array with only the required elements for the slice. 
	 This will reduce the memory used for the program. 
	 Syntax copy(dest, src)

	 The copy() function takes in two slices dest and src, and copies data from src to dest. 
	 It returns the number of elements copied.
	*/

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

	fmt.Printf("neededNumbers = %v\n", neededNumbers)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))

}

