package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var s string
	var first int
	var last int
	var length int

	fmt.Scan(&s)

	// input check
	if len(s) < 1 && 200000 < len(s) {

		fmt.Println("overflow")
		os.Exit(0)
	}

	s_array := strings.Split(s, "")
	first = strings.Index(s, "A")
	last = strings.LastIndex(s, "Z")

	length = len(s_array[first : last+1])

	fmt.Println(length)

}
