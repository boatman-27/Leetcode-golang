package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	current := head

	for current != nil {
		nextNode := current.Next
		current.Next = prev

		prev = current
		current = nextNode
	}

	return prev
}
