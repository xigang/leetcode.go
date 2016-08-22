package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var t *ListNode
	t = head.Next
	head.Next = swapPairs(head.Next.Next)
	t.Next = head
	return t
}
