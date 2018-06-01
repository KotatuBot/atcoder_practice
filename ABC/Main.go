package main

import "fmt"

func main() {

	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if 1 > a || b > 10000 {

		fmt.Println("Bad values")
	}

	sums := a * b

	if sums%2 == 0 {
		fmt.Println("Evn")
	} else {

		fmt.Println("Odd")
	}

}
