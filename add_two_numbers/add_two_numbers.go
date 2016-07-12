/**
 * You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 *  Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 *  Output: 7 -> 0 -> 8
**/

package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{}
	node := fakeHead
	node1 := l1
	node2 := l2
	hasCarry := false

	for node1 != nil || node2 != nil {
		node.Next = &ListNode{}
		node = node.Next

		val := 0
		if node1 != nil {
			val += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			val += node2.Val
			node2 = node2.Next
		}
		if hasCarry {
			val++
		}

		if val >= 10 {
			hasCarry = true
			val -= 10
		} else {
			hasCarry = false
		}
		node.Val = val
	}
	return fakeHead.Next
}
