package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	carryOver := 0
	for l1 != nil || l2 != nil || carryOver > 0 {
		sum := carryOver
		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		digitToStore := sum % 10
		carryOver = sum / 10

		current.Next = &ListNode{Val: digitToStore}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return result.Next
}
