package main

import (
	"fmt"
)

func main() {

	emails := map[string]string{"Job": "SWE", "Company": "xyz", "Salary": "great", "1": "1"}

	//Range with Map demo

	for v, m := range emails {

		fmt.Println(v, m)
		fmt.Printf("Values are %s = %s\n", v, m)
	}

}
