package main

import "fmt"

func main() {

	//Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	//Short Method
	for i := 1; i <= 10; i = i + 1 {
		fmt.Printf("Number %d\n", i)
	}

	//FizzBuzz
	for i := 1; i <= 45; i = i + 1 {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
