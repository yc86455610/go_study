package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2, l3 ListNode
	l1.Val = 2
	l1.Next = &l2
	l2.Val = 4
	l2.Next = &l3
	l3.Val = 3
	var ll1, ll2, ll3 ListNode
	ll1.Val = 5
	ll1.Next = &ll2
	ll2.Val = 6
	ll2.Next = &ll3
	ll3.Val = 4

	var r *ListNode
	r = addTwoNumbers(&l1, &ll1)
	fmt.Println(r.Val)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
