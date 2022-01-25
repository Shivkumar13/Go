package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter the range for fibonacci series")
	value := getInput()
	num, _ := strconv.Atoi(value)

	DisplayFibonacci(num)
}

func getInput() string {

	var data string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		data = scanner.Text()
	}

	return data
}

func DisplayFibonacci(num int) {

	n1 := 0
	n2 := 1

	fmt.Print(n1, ",", n2)

	for i := 0; i < num-2; i = i + 1 {
		t := n1 + n2
		n1 = n2
		n2 = t
		fmt.Print(",", t)
	}
	fmt.Printf("\n")
}
