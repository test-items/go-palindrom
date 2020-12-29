package main

import (
	"fmt"
)

func main() {
	inputString := "dsgsdfhsdggdsh"

	runeSlice := []rune(inputString)
	var result []rune
	for i := range runeSlice {
		for k := i + 1; k < len(runeSlice); k++ {
			if runeSlice[k] == runeSlice[i] {
				if findPalindrom(runeSlice[i : k+1]) {
					if len(runeSlice[i:k+1]) > len(result) {
						result = runeSlice[i : k+1]
					}
				}
			}
		}
	}
	fmt.Println(string(result))
}

func findPalindrom(runeSlice []rune) bool {
	k := len(runeSlice) - 1
	for i := 0; i <= k/2; i++ {
		if runeSlice[i] != runeSlice[k-i] {
			return false
		}
	}
	return true
}
