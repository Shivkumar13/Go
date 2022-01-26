package main

import (
	"fmt"
)

func main() {

	//fmt.Println("Hello enter range")
	fmt.Println(FiboRec(2))
}
func FiboRec(n int) int {

	if n == 0 || n == 1 {
		return n
	}
	return FiboRec(n-2) + FiboRec(n-1)

}
