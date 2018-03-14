package main

import "fmt"

func main() {
	a := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(swapPairs(&a))
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	p := head
	tail := head.Next.Next
	head = p.Next
	head.Next = p
	head.Next.Next = swapPairs(tail)
	return head
}
