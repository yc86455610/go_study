package main

func main() {
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	tmp := ListNode{0, head}
	p := &tmp
	q := &tmp
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return tmp.Next
}
