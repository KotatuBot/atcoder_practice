package main

import (
	"fmt"
	"strconv"
)

func main() {

	var value [3]int
	var test2 [3]int
	var test int

	for i := 0; i < 3; i++ {
		fmt.Scan(&value[i])
		test2[i] = strconv.Atoi(value[i])
	}

	test = test2[0]*100 + test2[1]*10 + test2[2]

	if test%4 == 0 {
		fmt.Println("YES")
	} else {

		fmt.Println("No")
	}

}
