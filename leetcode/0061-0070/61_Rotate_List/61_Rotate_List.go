package main

import "fmt"

func main() {
	a := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// a = ListNode{1, &ListNode{2, nil}}
	fmt.Println(rotateRight(&a, 2))
	a = ListNode{1, nil}
	fmt.Println(rotateRight(&a, 0))

}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l := 1
	p := head
	q := p.Next
	for q != nil {
		l++
		p = p.Next
		q = q.Next
	}
	p.Next = head
	k = l - k%l
	q = p.Next
	for i := 0; i < k; i++ {
		p = p.Next
		q = q.Next
	}
	p.Next = nil
	return q
}
