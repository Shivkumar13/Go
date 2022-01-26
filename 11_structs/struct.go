package main

import (
	"fmt"
)

type date struct {
	day   int
	month int
	year  int
}

func main() {

	var a date
	a = date{13, 7, 1997}

	fmt.Println(a.day)
	fmt.Println(a.month)
	fmt.Println(a.year)
}
