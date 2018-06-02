package main

import "fmt"

func main() {

	var group int
	var i int = 0
	var lower [100100]int
	var uper [100100]int
	var score int = 0

	// get group
	fmt.Scan(&group)

	// get lower - uper+1
	for ; i < group; i++ {
		fmt.Scan(&lower[i])
		fmt.Scan(&uper[i])
		score += uper[i] - lower[i] + 1

	}
	// score
	fmt.Println(score)

}
