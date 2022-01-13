package main

import "fmt"

func main() {
	//Arrays
	var OrgArr [2]string

	OrgArr[0] = "Apple"
	OrgArr[1] = "Google"

	//Array with initialization
	var fruitArr [2]string
	fruitArr = [2]string{"apple", "orange"}

	//	OrgArr[2] = "Red Hat"

	fmt.Println(OrgArr)
	fmt.Println(fruitArr)
}
