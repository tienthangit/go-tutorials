/*
Syntax:
		var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
		b := map[KeyType]ValueType{key1:value1, key2:value2,...}

*/

package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)

	/* Output
	a   map[brand:Ford model:Mustang year:1964]
	b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]
	*/


	/* Using make

	var a = make(map[KeyType]ValueType)
	b := make(map[KeyType]ValueType)

	*/

	var c = make(map[string]string) // The map is empty now
  c["brand"] = "Ford"
  c["model"] = "Mustang"
  c["year"] = "1964"
                                
  var d = make(map[string]int)
  d["Oslo"] = 1
  d["Bergen"] = 2
  d["Trondheim"] = 3
  d["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)

	/* Output:
		a   map[brand:Ford model:Mustang year:1964]
		b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]
	*/

	
	/*
		Remove Element from Map
		delete(map_name, key)
	*/

	var e = make(map[string]string)
  e["brand"] = "Ford"
  e["model"] = "Mustang"
  e["year"] = "1964"

  fmt.Println(a)

	/* Output
		map[brand:Ford model:Mustang year:1964]
		map[brand:Ford model:Mustang]
	*/

	/* Check For Specific Elements in a Map
		val, ok :=map_name[key]
	*/

	var f = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

  val1, ok1 := f["brand"] // Checking for existing key and its value
  val2, ok2 := f["color"] // Checking for non-existing key and its value
  val3, ok3 := f["day"]   // Checking for existing key and its value
  _, ok4 := f["model"]    // Only checking for existing key and not its value

  fmt.Println(val1, ok1)
  fmt.Println(val2, ok2)
  fmt.Println(val3, ok3)
  fmt.Println(ok4)

	/* Output
		Ford true
		false
		true
		true
	*/
}



