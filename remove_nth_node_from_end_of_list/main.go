package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Given linked list: 1->2->3->4->5, and n = 2.
//After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:
// Given n will always be valid.
// Try to do this in one pass.

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{-1, head}
	p := &dummy
	q := &dummy

	for i := 0; i < n; i++ { //先抵达n
		q = q.Next
	}

	for q.Next != nil { //一起走
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next
	return dummy.Next
}
