package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Solution with recursion
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	finalNode := &ListNode{}

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{Val: 0}
	}

	if l2 == nil {
		l2 = &ListNode{Val: 0}
	}

	if sum := l1.Val + l2.Val; sum > 9 {
		finalNode.Val = sum % 10

		switch {
		case l1.Next != nil && l2.Next != nil:
			l1.Next.Val++
			finalNode.Next = addTwoNumbers(l1.Next, l2.Next)

		case l1.Next != nil && l2.Next == nil:
			l1.Next.Val++
			finalNode.Next = addTwoNumbers(l1.Next, l2.Next)

		case l1.Next == nil && l2.Next != nil:
			l2.Next.Val++
			finalNode.Next = addTwoNumbers(l1.Next, l2.Next)

		default:
			finalNode.Next = &ListNode{
				Val: 1,
			}
		}

		return finalNode
	}

	finalNode.Val = l1.Val + l2.Val
	finalNode.Next = addTwoNumbers(l1.Next, l2.Next)

	return finalNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(*addTwoNumbers(l1, l2))
}
