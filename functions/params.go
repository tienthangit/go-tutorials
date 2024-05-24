/*
	func FunctionName(param1 type, param2 type, param3 type) {
		// code to be executed
	}


*/

package main

import (
	"fmt"
)

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func familyName2(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func _main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")

	familyName2("Liam", 3)
	/*
	Hello Liam Refsnes
	Hello Jenny Refsnes
	Hello Anja Refsnes
	*/
}