package strings

import (
	"fmt"
	"strings"
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
		fmt.Println("Longest substring in ", one[i], " is how many long without repeated characters? ", lengthOfLongestSubstring(one[i]))
		fmt.Println("Longest substring in ", two[i], " is how many long without repeated characters? ", lengthOfLongestSubstring(two[i]))
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

func lengthOfLongestSubstring(s string) int {
    strSlice := strings.Split(s,"")
    countOcc := map[string] int {}
    length := 0
    longest := 0
    for i,literal := range strSlice{
        if val, ok := countOcc[literal];ok{
            diff := i- val
            if length >= diff {
                length = diff -1 
            }
        }
        countOcc[literal] = i
        length ++
        if longest < length{
                longest = length
        }
    }
    return longest
    
}
