package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	_, err := fmt.Scan(&n, &m, &k)
	if err != nil {
		return
	}

	minutesRequired := n * k

	timeSpent := minutesRequired / m

	if minutesRequired%m != 0 {
		timeSpent++
	}

	fmt.Println(timeSpent)
}
