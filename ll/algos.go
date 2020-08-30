package ll

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Algos() {
	l1 := &ListNode{Val: 8, Next: &ListNode{9, &ListNode{9, nil}}}
	l2 := &ListNode{Val: 2, Next: nil}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println("Sum of number: ", l1.Next.Next.Val, l1.Next.Val, l1.Val, "and number: ", l2.Val, "is: ", l3.Next.Next.Next.Val, l3.Next.Next.Val, l3.Next.Val, l3.Val)
}

/**
* Numbers are stored in reverse order
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == l2 && l1 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carryover, v3 := addOneByOne(l1.Val, l2.Val, 0)

	var l3 *ListNode
	l3 = &ListNode{Val: v3, Next: nil}
	head := l3
	for {
		if l1.Next == l2.Next && l1.Next == nil {
			if carryover != 0 {
				l3.Next = &ListNode{Val: carryover, Next: nil}
			}
			break
		}
		if l2.Next == nil {
			l1 = l1.Next
			carryover, v3 = addOneByOne(l1.Val, carryover, 0)
			l3.Next = &ListNode{Val: v3, Next: nil}
			l3 = l3.Next

			if l1.Next == nil {
				if carryover != 0 {
					l3.Next = &ListNode{Val: carryover, Next: nil}
				}
				break
			}
		} else if l1.Next == nil {
			l2 = l2.Next
			carryover, v3 = addOneByOne(l2.Val, carryover, 0)
			l3.Next = &ListNode{Val: v3, Next: nil}
			l3 = l3.Next

			if l2.Next == nil {
				if carryover != 0 {
					l3.Next = &ListNode{Val: carryover, Next: nil}
				}
				break
			}

		} else {
			l1 = l1.Next
			l2 = l2.Next
			carryover, v3 = addOneByOne(l1.Val, l2.Val, carryover)
			l3.Next = &ListNode{Val: v3, Next: nil}
			l3 = l3.Next
		}
	}
	return head
}

func addOneByOne(v1 int, v2 int, carryover int) (int, int) {
	val := v1 + v2 + carryover
	if val > 9 {
		val = val % 10
		return 1, val
	}
	return 0, val
}
