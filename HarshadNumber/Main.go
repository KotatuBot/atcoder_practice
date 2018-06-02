package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func harshad_number(number int) {

	var sums int = 0
	number_str := strconv.Itoa(number)
	array := strings.Split(number_str, "")

	for i := 0; i < len(array); i++ {

		transfer, _ := strconv.Atoi(array[i])
		sums += transfer

	}

	if number%sums == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func main() {

	var x int

	fmt.Scan(&x)

	if x < 1 && 100000000 < x {

		fmt.Println("Overflow")
		os.Exit(0)
	}

	harshad_number(x)

}
