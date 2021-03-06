package slices

import (
	"fmt"
	"strings"
)

func Algos() {

	list := []int{1, 3, 5, 6, 7, 90}
	sum := 10
	fmt.Println("Two numbers exist in the given list:", list, "with the given sum: ", sum, " is ", findTwoNumbersWhoseSumIsEqualToAGivenValue(list, sum))

	list = []int{1, 1, 3, 4, 6, 6, 8, 90}
	fmt.Println("List with Duplicates: ", list, "List without duplicates: ", removeDuplicatesFromASortedList(list))
	list = []int{1, 2, 4, 6, 7, 8, 10}
	fmt.Println("Does Pythagorean Triplet for ", list, "exists? ", findPythagoreanTriplet(list))
	list = []int{7, 1, 18, 6, 5, 89, 14}
	fmt.Println("Does Pythagorean Triplet for ", list, "exists? ", findPythagoreanTriplet(list))
}

/*

Find any two numbers in an ordered slice or a list, whose sum is equal to given number

Assumptions:

Negative numbers allowed,
No fractions allowed; only integers,
Numbers in the list can be greater than or less than given number
Repeat numbers allowed,
Can not add a number with itself unless it is repeated

*/

func findTwoNumbersWhoseSumIsEqualToAGivenValue(numList []int, sum int) bool {

	i := 0
	j := len(numList) - 1 //Last element as first element starts from 0

	for {
		if i == j {
			break
		}
		firstNum := numList[i]
		secondNum := numList[j]
		if firstNum+secondNum > sum {
			j = j - 1
		} else if firstNum+secondNum < sum {
			i = i + 1
		} else {
			return true
		}
	}
	return false
}

/*

Remove duplicates from a sorted list
Assumptions:

Negative numbers allowed,
No fractions allowed; only integers,
List is Sorted
Only printing of the distinct list without duplicates is needed.
If there is a need to persist the changes pass the slice as a pointer
or return a list if in place updates are not needed

*/

func removeDuplicatesFromASortedList(numList []int) []int {
	var expectedList []int
	for i := 0; i < len(numList)-1; i++ {
		if numList[i] == numList[i+1] {
			continue
		}
		expectedList = append(expectedList, numList[i])
	}
	expectedList = append(expectedList, numList[len(numList)-1])
	return expectedList
}

/* Check if a pythagorean triplet exists

Assumptions:
1. All numbers are integers (positive)
2. Square of all numbers is less than max limit of the type

*/

func findPythagoreanTriplet(numbers []int) bool {
	squareToRoot := make(map[int]bool)

	for i := 0; i < len(numbers); i++ {
		square := numbers[i] * numbers[i]
		squareToRoot[square] = true
	}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			sumOfSquares := numbers[i]*numbers[i] + numbers[j]*numbers[j]
			// if key not present golang map implementation will return false
			if squareToRoot[sumOfSquares] {
				return true
			}
		}
	}
	return false

}

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.


Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

func letterCombinations(digits string) []string {

	dToStr := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	digitsSlice := strings.Split(digits, "")
	intermediate := []string{}
	for _, val := range digitsSlice {
		if val == "1" {
			continue
		}
		if len(intermediate) == 0 {
			intermediate = dToStr[val]
		} else {
			newSlice := []string{}
			for _, letter := range dToStr[val] {
				for _, seq := range intermediate {
					newSlice = append(newSlice, seq+letter)
				}
			}
			intermediate = newSlice
		}
	}
	return intermediate
}

/*

Notes:

1. Printing values of different types in fmt.Println can be accomplished by using commas instead of + which would expect all values to be strings
2. append function definition is append( slice_tobe_appended_to, values_tobe_appended). This is function that works for all standard types with expectation that both slice and value are of same type
3.
*/
