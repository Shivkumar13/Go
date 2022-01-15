package main

import "fmt"

func main() {

	numbers := []int{2, 120, 20, 1, 1}

	// Loop through the numbers

	for _, ID := range numbers {
		fmt.Printf("The numbers in the array serially are ID:%d\n", ID)
	}

	wholesum := 0

	for _, ele := range numbers {

		wholesum += ele // wholesum = wholesum + ele

	}
	fmt.Println(wholesum)
}
