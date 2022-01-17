package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 3, 7, 6, 8, 9, 10}
	fmt.Println(getConcatenation(nums))

}

func getConcatenation(nums []int) []int {

	var ans []int
	for i := 0; i <= 1; i = i + 1 {
		for _, e := range nums {
			ans = append(ans, e)
		}
	}
	return ans
}
