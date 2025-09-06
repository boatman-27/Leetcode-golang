package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// find length
	current := head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}

	if length == n {
		return head.Next
	}

	index := length - n - 1
	node := head
	for index > 0 {
		node = node.Next
		index--
	}

	toRemove := node.Next
	node.Next = toRemove.Next

	return head
}
