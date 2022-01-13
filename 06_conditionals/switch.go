package main

import (
	"fmt"
)

func main() {

	var color string
	fmt.Println("Enter an input to SWITCH Block 1:")
	fmt.Scanf("%v", &color)

	var hi int
	fmt.Println("Enter an input to SWITCH Block 2:")
	fmt.Scanf("%v", &hi)

	switch color {

	case "red":
		fmt.Println("This is red")
	case "green":
		fmt.Println("This is Green")
	case "hi":
		fmt.Println("This is NOT color")
	default:
		fmt.Println("This is Nothing")
	}

	switch hi {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("This is 2")
	default:
		fmt.Println("It is default statement.")
	}
}
