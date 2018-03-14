package main

import (
	"fmt"
)

func main() {
	a := ListNode{2, nil}
	b := ListNode{1, nil}
	c := mergeTwoLists(&a, &b)
	fmt.Println(c)
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	head := l1

	for l1 != nil && l2 != nil {
		if l2.Val >= l1.Val && (l1.Next == nil || l2.Val < l1.Next.Val) {
			tmp := l2
			l2 = l2.Next
			tmp.Next = l1.Next
			l1.Next = tmp
		} else {
			l1 = l1.Next
		}
	}
	if l2 != nil {
		l1.Next = l2
	}

	return head
}
