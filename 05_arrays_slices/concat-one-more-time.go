package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter len of an Array you want to enter")

	fmt.Scanf("%d", &n)

	input := make([]int, n)

	for i := 0; i < n; i = i + 1 {
		fmt.Printf("Enter %dth elements in the INPUT Array", i)
		fmt.Scanf("%d", &input[i])
	}

	fmt.Println("Your Array is: ", input)

	getForMe(input)

}

func getForMe(input []int) []int {

	var res []int

	for i := 0; i < 2; i++ {
		for _, ele := range input {

			res = append(res, ele)

		}
	}

	fmt.Println("Concatenated Array is: ", res)

	return res
}
