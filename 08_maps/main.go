package main

import "fmt"

func main() {
	//Define Map
	var somemap map[string]string
	somemap = make(map[string]string)

	//Adding Key-value to the map
	somemap["shivkumar"] = "Ople"
	somemap["Amey"] = "NickAmey"
	somemap["Ambuja"] = "Cement"

	fmt.Println(somemap)
	fmt.Println(len(somemap))

	//Delete from a Map
	delete(somemap, "shivkumar") /* Case sensitive behavior being shown when deleting a key from map
	   As the key shivkumar's `s` should be as it is. */
	fmt.Println(somemap["Amey"])
	fmt.Println(somemap)
}
