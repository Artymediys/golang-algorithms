package main

import (
	"fmt"
	"math"
)

func minGoodSubstringLength(n int, s string) int {
	if len(s) != n {
		return -1
	}

	left, minLength := 0, math.MaxInt32
	charCount := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0}

	for right, char := range s {
		if _, ok := charCount[char]; ok {
			charCount[char]++
		}

		for charCount['a'] > 0 && charCount['b'] > 0 && charCount['c'] > 0 && charCount['d'] > 0 {
			minLength = int(math.Min(float64(minLength), float64(right-left+1)))
			if _, ok := charCount[rune(s[left])]; ok {
				charCount[rune(s[left])]--
			}
			left++
		}
	}

	if minLength == math.MaxInt32 {
		return -1
	}
	return minLength
}

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	minLength := minGoodSubstringLength(n, s)
	fmt.Println(minLength)
}
