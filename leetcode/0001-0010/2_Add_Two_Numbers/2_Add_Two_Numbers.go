package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head ListNode
	var tmp ListNode
	var l = &head
	for l1 != nil || l2 != nil {
		var ln ListNode
		l.Next = &ln
		l = l.Next
		if l1 == nil {
			l1 = &tmp
		}
		if l2 == nil {
			l2 = &tmp
		}
		sum := l1.Val + l2.Val + carry
		l.Val, carry = sum%10, sum/10
		l1, l2 = l1.Next, l2.Next
	}
	if carry != 0 {
		var ln ListNode
		l.Next = &ln
		l = l.Next
		l.Val += carry
	}
	return head.Next
}

func main() {

}
