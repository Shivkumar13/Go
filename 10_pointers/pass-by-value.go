package main

import "fmt"

var p int = 10

//var q int = 10

var q *int = &p

func main() {

	add(-3, -5)
	fmt.Println(*q)
}

func add(p int32, q int32) {

	c := p + q
	fmt.Println("Addition is:", c)
	//	return c
}
