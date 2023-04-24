package main

import "fmt"

func romanToInt(romanNum string) int {
	charMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	romanNumLength := len(romanNum)

	if romanNumLength == 0 {
		return 0
	} else if romanNumLength == 1 {
		return charMap[romanNum[0]]
	}

	resultInt := charMap[romanNum[romanNumLength-1]]

	for i := romanNumLength - 2; i >= 0; i-- {
		if charMap[romanNum[i]] < charMap[romanNum[i+1]] {
			resultInt -= charMap[romanNum[i]]
		} else {
			resultInt += charMap[romanNum[i]]
		}
	}

	return resultInt
}

func main() {
	fmt.Println(romanToInt("XII"))
}
