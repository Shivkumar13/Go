package main

import "fmt"

func main() {

	c := 4
	a := &c
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T \t %T\n", a, b)

	//Use * to read the value from address

	fmt.Println(*b) // Showing value stored in `a` which is again and address of `c`

	fmt.Println(**b) // So we do **b, which will show the `stored int value`

	fmt.Println(*a)  // Grabbing out value from the address.
	fmt.Println(*&c) // This is same as `*a`, because `a := &c`

}
