package main

import (
	"fmt"
)

func main() {
	var h1, h2, h3, h4 int
	_, err := fmt.Scanf("%d %d %d %d", &h1, &h2, &h3, &h4)
	if err != nil {
		return
	}

	if (h1 <= h2 && h2 <= h3 && h3 <= h4) || (h1 >= h2 && h2 >= h3 && h3 >= h4) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
