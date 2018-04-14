package main

import (
	"fmt"
)

func main() {
	test(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	// test(&ListNode{1, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}, &ListNode{1, &ListNode{4, nil}})
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
	res := &ListNode{0, head}
	newHead := res

	p := head
	sign := true
	q := head.Next
	for q != nil {
		if p.Val == q.Val {
			sign = false
		} else {
			if sign {
				newHead.Next = p
				p = q
				sign = true
				newHead = newHead.Next
			} else {
				p = q
				sign = true
			}
		}
		q = q.Next
	}
	if sign {
		newHead.Next = p
		newHead = newHead.Next
	}
	newHead.Next = nil

	return res.Next
}
