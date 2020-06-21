package slices

import (
  "fmt"
//  "strings"
)


func Algos(){

  list := []int{1, 3, 5, 6, 7, 90}
  sum := 10
  fmt.Println("Two number exist in the given list with the given sum: ", findTwoNumbersWhoseSumIsEqualToAGivenValue(list, sum)) 
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

func findTwoNumbersWhoseSumIsEqualToAGivenValue(numList []int, sum int)(bool){

  i := 0
  j := len(numList) - 1 //Last element as first element starts from 0

  for {
    if i == j{
      break;
    }
    firstNum := numList[i]
    secondNum := numList[j]
    if firstNum + secondNum > sum {
       j = j - 1
    }else if firstNum + secondNum < sum {
       i = i + 1
    }else {
      return true
    } 
  } 
  return false 
}
