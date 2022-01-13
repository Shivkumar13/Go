package main

import "fmt"

func main() {

	x := 5
	y := 10

	//If else
	if x <= y {

		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Println("%d is less than %d\n", y, x)
	}

	// else if
	color := "green"

	if color == "red" {
		fmt.Println("color is red")

	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

}
