package main

import (
	"fmt"
)

func Greetings(name string) string {

	return "Hello " + name
}

func main() {
	fmt.Println(Greetings("Shivkumar"))
}
