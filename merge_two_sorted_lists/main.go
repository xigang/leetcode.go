package main

// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		temp := l2
		temp.Next = mergeTwoLists(l1, l2.Next)
		return temp
	} else {
		temp := l1
		temp.Next = mergeTwoLists(l1.Next, l2)
		return temp
	}
}
