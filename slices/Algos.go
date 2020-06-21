package slices

import (
	"fmt"
	//  "strings"
)

func Algos() {

	list := []int{1, 3, 5, 6, 7, 90}
	sum := 10
	fmt.Println("Two numbers exist in the given list:", list, "with the given sum: ", sum, " is ", findTwoNumbersWhoseSumIsEqualToAGivenValue(list, sum))

	list = []int{1, 1, 3, 4, 6, 6, 8, 90}
	fmt.Println("List with Duplicates: ", list, "List without duplicates: ", removeDuplicatesFromASortedList(list))
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

Find any two numbers in an ordered slice or a list, whose sum is equal to given number

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

/*

Notes:

1. Printing values of different types in fmt.Println can be accomplished by using commas instead of + which would expect all values to be strings
2. append function definition is append( slice_tobe_appended_to, values_tobe_appended). This is function that works for all standard types with expectation that both slice and value are of same type

*/
