package main

import (
	"fmt"
	"os"
)

func sum(s []int) int {

	var sums int = 0

	for i := 0; i < len(s); i++ {
		sums += s[i]
	}
	return sums

}

func min(j []int) int {

	var min_value int = j[0]

	for v := 1; v < len(j); v++ {

		if j[v] < min_value {
			min_value = j[v]

		}
	}

	return min_value

}

func main() {

	var types int
	var max int
	var i int = 0
	var pauder [10010]int
	var rest int
	var result int

	fmt.Scan(&types)
	fmt.Scan(&max)

	if types < 2 && 1000 < types {
		fmt.Println(" types is 2 < types < 1000")
		os.Exit(0)

	}

	if 100000 < max {
		fmt.Println("overflow")
		os.Exit(0)

	}

	for ; i < types; i++ {
		fmt.Scan(&pauder[i])

	}

	pauder_arry := pauder[:types]

	sums := sum(pauder_arry)
	min_value := min(pauder_arry)

	// rest
	rest = max - sums

	// rest
	rest_value := rest / min_value

	result = types + rest_value

	fmt.Println(result)

}
