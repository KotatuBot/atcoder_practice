package main

import "fmt"

func main() {

	var one int
	var two int
	var three int
	var name string

	fmt.Scan(&one)
	fmt.Scan(&two)
	fmt.Scan(&three)
	fmt.Scan(&name)

	sum := one + two + three

	fmt.Println(sum, name)

}
