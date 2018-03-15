package main

import "fmt"

func main() {
	a := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(reverseKGroup(&a, 2))
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	p := head
	nHead := &ListNode{0, nil}
	for i := 0; i < k; i++ {
		tmp := nHead.Next
		nHead.Next = p
		p = p.Next
		nHead.Next.Next = tmp
	}
	head.Next = reverseKGroup(tail, k)
	return nHead.Next
}
