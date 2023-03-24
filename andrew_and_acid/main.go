package main

import "fmt"

func IsPossibleToEqualize(slice []uint64) bool {
	var value uint64 = slice[0]

	for i := 1; i < len(slice); i++ {
		if value > slice[i] {
			return false
		}

		value = slice[i]
	}

	return true
}

func FirstMaxElement(slice []uint64) (int, uint64) {
	var maxID int = 0
	var maxValue uint64 = slice[0]

	if len(slice) == 1 {
		return maxID, maxValue
	}

	for i := 1; i < len(slice); i++ {
		if maxValue < slice[i] {
			maxID = i
			maxValue = slice[i]
		}
	}
	return maxID, maxValue
}

func main() {
	// Entering the count of acid tanks
	var n uint32
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Entering the number of liters in each acid tank
	a := make([]uint64, n)
	for i := uint32(0); i < n; i++ {
		_, err = fmt.Scan(&a[i])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Check for the possibility of acid tank equation
	if !IsPossibleToEqualize(a) {
		fmt.Println(-1)
		return
	}

	// Count the min operations
	var minOperations uint64 = 0
	maxValID, maxVal := FirstMaxElement(a)
	for k := maxValID; k > 0; {
		for i := k - 1; i >= 0; i-- {
			a[i]++
			if a[i] == maxVal {
				k--
			}
		}
		minOperations++
	}

	fmt.Println(minOperations)
}
