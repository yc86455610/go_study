package main

import (
	"fmt"
)

func main() {
	// test(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	test(&ListNode{1, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	// test(&ListNode{1, &ListNode{3, &ListNode{3, &ListNode{3, nil}}}}, &ListNode{1, nil})

}

func test(head *ListNode, r *ListNode) {
	head = deleteDuplicates(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	q := head.Next
	for q != nil {
		if p.Val != q.Val {
			p.Next = q
			p = p.Next
		}
		q = q.Next
	}
	p.Next = nil

	return head
}
