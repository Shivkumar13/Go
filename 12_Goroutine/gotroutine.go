package main

import (
	"fmt"
	"time"
)

//A Goroutine is a function or method which executes independantly and simultaneously in connection with
//any other Goroutines  present in your program.

//Or in other words every concurrently executing activity in Go languages is known as Goroutines.

//You can consider Goroutine like a lighetweight thread.

func display(str string) {

	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	go display("Welcome")
	display("Geeks from India")
}
