package main

import "fmt"

func main() {
	a := []*ListNode{&ListNode{0, nil}, &ListNode{1, nil}}
	fmt.Println(mergeKLists(a))

}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 0 {
		return nil
	}
	a := mergeKLists(lists[:len(lists)/2])
	b := mergeKLists(lists[len(lists)/2:])
	return mergeTwoLists(a, b)
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
