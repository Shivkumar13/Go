package main

import "fmt"

func main() {

	numbers := []int{2, 120, 20, 1, 1}

	// Loop through the numbers

	for _, ID := range numbers {
		fmt.Printf("The numbers in the array serially are ID:%d\n", ID)
	}

	addition := 0

	fmt.Println(addition)

	for _, something := range numbers {
		addition += something

		fmt.Println(addition)

		fmt.Printf("This is something %d\n", something)
	}
	fmt.Println("....")
	fmt.Println("...")
	fmt.Println(addition)
}
