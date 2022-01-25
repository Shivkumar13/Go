package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter the range of Fibonacci")

	value := getSome()
	num, _ := strconv.Atoi(value)

	DisplaySeries(num)

}

func getSome() string {

	var data string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {

		data = scanner.Text()
	}

	return data

}

func DisplaySeries(num int) {

	num_1 := 0
	num_2 := 1

	fmt.Print(num_1, ",", num_2)
	for i := 0; i < num-2; i = i + 1 {
		add := num_1 + num_2
		num_1 = num_2
		num_2 = add

		fmt.Print(",", add)

	}
}
