package main

import (
	"fmt"
)

func Greetings(name string) string {

	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2

}
func main() {
	fmt.Println(Greetings("Shivkumar"))
	fmt.Println(getSum(4, 5))

}
