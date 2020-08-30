package stacks

import (
	"fmt"
	"strings"
)

func Algos() {

	pListofList := [][]string{
		{"(", ")", "(", "(", "(", ")"},
		{"(", ")", "(", "(", "(", ")", ")"},
		{")", ")", ")", "(", "("},
		{"(", ")", ")", ")", "(", "(", "(", ")", "(", ")", ")", "(", ")"},
		{"(", "(", ")", "(", ")", ")"},
		{"(", "(", "(", ")", "(", ")", ")"},
		{"(", "(", "(", ")", ")", ")", ")"},
		{"(", ")", ")", ")", ")", "(", "("},
	}
	for _, p := range pListofList {
		fmt.Println("Length of longest continuous balanced paranthesis sequence: ", longestValidParentheses(strings.Join(p, "")), "in string", strings.Join(p, ""))
	}
}

type stack struct {
	content []int
}

func (s *stack) push(p int) {
	//... allows append of two int slices as variadic arguments are supported
	s.content = append([]int{p}, s.content...)
}

// replaces stack by removing first element to follow LIFO
func (s *stack) pop() {
	s.content = s.content[1:]
}

func (s *stack) isEmpty() bool {

	if len(s.content) == 0 {
		return true
	}
	return false
}

// find the value for last element
func (s *stack) peek() int {
	return s.content[0]
}

func max(v1 int, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func longestValidParentheses(p string) int {

	var (
		l = 0
		s = &stack{
			content: []int{-1},
		} //this covers the scenario where the entire string is perfectly balanced
	)

	for i, c := range p {
		if c == '(' {
			s.push(i) //keep pushing with hopes for getting the other paranthesis
			continue
		}
		s.pop() //pop with intent to find a pair
		if s.isEmpty() {
			s.push(i)
			continue
		}
		l = max(l, i-s.peek()) //allows us to change length based on total contiguous balanced pairs

	}
	return l
}
