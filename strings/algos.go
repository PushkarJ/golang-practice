package strings

import (
	"fmt"
	"strconv"
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
	for i := 1; i < 30; i++ {
		fmt.Println("Count and say for iteration ", strconv.Itoa(i), "is: ", countAndSay(i))
	}
	for i := 0; i < len(one); i++ {
		fmt.Println("Are ", one[i], " and ", two[i], " strings anagrams?", areTwoStringsAnagrams(one[i], two[i]))
		fmt.Println("Longest substring in ", one[i], " is how many long without repeated characters? ", lengthOfLongestSubstring(one[i]))
		fmt.Println("Longest substring in ", two[i], " is how many long without repeated characters? ", lengthOfLongestSubstring(two[i]))
	}
	sboard := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaab"}
	word := "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	var board [][]byte
	for _, s := range sboard {
		board = append(board, []byte(s))
	}
	fmt.Println("Word Exists on the board: ", exist(board, word))

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
	strSlice := strings.Split(s, "")
	countOcc := map[string]int{}
	length := 0
	longest := 0
	for i, literal := range strSlice {
		if val, ok := countOcc[literal]; ok {
			diff := i - val
			if length >= diff {
				length = diff - 1
			}
		}
		countOcc[literal] = i
		length++
		if longest < length {
			longest = length
		}
	}
	return longest

}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	value := countAndSay(n - 1)
	output := ""
	ele := ""
	count := 0
	for _, c := range strings.Split(value, "") {
		if ele == "" {
			ele = c
			count++
		} else if ele == c {
			count++
		} else {
			cs := strconv.Itoa(count)
			output += cs + ele
			ele = c
			count = 1
		}
	}
	cs := strconv.Itoa(count)
	output += cs + ele

	return output

}

// NOTE: If you want non-repeating sequence, replace the visited character with something out of range
/*
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.

*/
func exist(board [][]byte, word string) bool {
	wb := []byte(word)
	found := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			//fmt.Println(board[i][j], wb[0])

			if board[i][j] == wb[0] {
				//fmt.Println(board[i][j], wb[0])
				found = findNbr(board, i, j, wb, 1)
				if found {
					return true
				}
			}
		}
	}
	return found
}

func findNbr(board [][]byte, i int, j int, word []byte, c int) bool {
	found := false
	l := board[i][j]
	board[i][j] = byte('@')
	if len(word) <= c {
		//fmt.Println("reached")
		return true
	}
	if len(board) <= i && len(board[i]) <= j {
		return false
	}

	if len(board) > i+1 && board[i+1][j] == word[c] {
		//fmt.Println(i+1, j, c)
		rfound := findNbr(board, i+1, j, word, c+1)
		if found == false {
			found = rfound
		}

	}
	if len(board[i]) > j+1 && board[i][j+1] == word[c] {
		//fmt.Println(i, j+1, c)
		dfound := findNbr(board, i, j+1, word, c+1)
		if found == false {
			found = dfound
		}
	}

	if len(board) > i && i > 0 && board[i-1][j] == word[c] {
		//fmt.Println(i-1, j, c)
		lfound := findNbr(board, i-1, j, word, c+1)
		if found == false {
			found = lfound
		}

	}
	if len(board[i]) > j && j > 0 && board[i][j-1] == word[c] {
		//fmt.Println(i, j-1, c)
		ufound := findNbr(board, i, j-1, word, c+1)

		if found == false {
			found = ufound

		}
	}
	if found == false {
		board[i][j] = l
	}
	return found
}
