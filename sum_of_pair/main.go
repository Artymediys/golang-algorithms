package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Finding any pair
func findPair(A []int32, k int32) []int {
	// Map init
	searchMap := make(map[int32]int, len(A))
	for index, value := range A {
		searchMap[value] = index
	}

	// Searching
	for key, value := range searchMap {
		if element, ok := searchMap[k-key]; ok {
			return []int{value, element}
		}
	}

	return []int{}
}

func main() {
	// Setting random range and seed
	var minRand, maxRand int32 = 1, 10
	rand.NewSource(time.Now().Unix())

	// Creating the slice with a random len/cap and elements
	A := make([]int32, rand.Int31n(maxRand-minRand)+minRand)
	for i := 0; i < len(A); i++ {
		A[i] = rand.Int31n(maxRand-minRand) + minRand
	}

	// Creating the sum-variable with a random value
	k := rand.Int31n(maxRand-minRand) + minRand

	fmt.Printf("A: %v, k: %d\n", A, k)
	fmt.Println("Pair's indexes:", findPair(A, k))
}
