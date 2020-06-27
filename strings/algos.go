package strings

import (
	"fmt"
)

func Algos() {

	one := []string{"silent", "bay area", "tenet"}
	two := []string{"listen", "san francisco", "tenet"}
	//all palindromes are anagrams but not all anagrams are palindromes
	//synonyms != anagrams

	if len(one) != len(two) {
		return
	}

	for i := 0; i < len(one); i++ {
		fmt.Println("Are ", one[i], " and ", two[i], " strings anagrams?", areTwoStringsAnagrams(one[i], two[i]))
	}

}

func areTwoStringsAnagrams(one string, two string) bool {

	checkIfPresent := map[rune]int{}
	//assuming only 26 alphabets in scope. Below string can be updated to include more such as numbers
	universeOfStringChars := "abcdefghijklmnopqrstuvwxyz"

	for _, c := range universeOfStringChars {
		checkIfPresent[c] = 0
	}

	for _, c := range one {
		checkIfPresent[c] = 1
	}

	for _, c := range two {
		checkIfPresent[c] = 0
	}

	for _, c := range universeOfStringChars {
		if checkIfPresent[c] != 0 {
			return false
		}
	}
	return true
}
