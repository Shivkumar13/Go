package main

import (
	"fmt"
)

func main() {

	review := map[string]int{"Nike": 1, "Adidas": 2, "Puma": 3}

	//Just sorting the keys by using range

	for i, k := range review { // printing sequence is changing
		fmt.Printf("Listed Brands are %s = %d\n", i, k)
	}
}
