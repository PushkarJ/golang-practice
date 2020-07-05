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
	list = []int{1, 2, 4, 6, 7, 8, 10}
	fmt.Println("Does Pythagorean Triplet for ", list, "exists? ", findPythagoreanTriplet(list))
	list = []int{7, 1, 18, 6, 5, 89, 14}
	fmt.Println("Does Pythagorean Triplet for ", list, "exists? ", findPythagoreanTriplet(list))
	pList := []string{"(", ")", "(", "(", "(", ")"}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{"(", ")", "(", "(", "(", ")", ")"}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{")", ")", ")", "(", "("}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{"(", ")", ")", ")", "(", "(", "(", ")", "(", ")", ")", "(", ")"}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{"(", "(", ")", "(", ")", ")"}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{"(", "(", "(", ")", "(", ")", ")"}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{"(", "(", "(", ")", ")", ")", ")"}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
	pList = []string{"(", ")", ")", ")", ")", "(", "("}
	fmt.Println("Length of longest continuous balanced paranthesis sequence: ", lengthOfLongestBalancedParanthesisString(pList))
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

/*

Find the longest continuous sequence of balanced paranthesis.

Assumptions: String contains only paranthesis of one type e.g. ( )

*/

func lengthOfLongestBalancedParanthesisString(paranthesisList []string) int {
	left := 0
	right := 0
	length := 0
	previous := paranthesisList[0]
	if previous == "(" {
		left = 1
	} //if right paranthesis is found that is anyway unbalanced so no need to increment
	for i := 1; i < len(paranthesisList); i++ {
		if paranthesisList[i] == "(" {
			left++
		} else if paranthesisList[i] == ")" && left > 0 {
			//there has to be atleast one left paranthesis for possibility of a balanced list
			right++
		}
		if left < right {
			if length < left*2 {
				length = left * 2 //update length only if it is higher than current length
			}
			//when left paranthesis becomes zero any number of right paranthesis before it need to be reset
			if previous == paranthesisList[i] {
				left = 0
				right = 0
			}
		} else if left >= right {
			if length < right*2 {
				length = right * 2
			}
			//reset right only if there are consecutive left paranthesis
			if previous == paranthesisList[i] && previous != ")" {
				right = 0
			}
		}
		previous = paranthesisList[i]
	}
	return length
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

Notes:

1. Printing values of different types in fmt.Println can be accomplished by using commas instead of + which would expect all values to be strings
2. append function definition is append( slice_tobe_appended_to, values_tobe_appended). This is function that works for all standard types with expectation that both slice and value are of same type
3.
*/
