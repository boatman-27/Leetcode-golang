package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	set := make(map[int]bool)
	set[head.Val] = true

	current := head
	for current != nil && current.Next != nil {
		if set[current.Next.Val] {
			current.Next = current.Next.Next
		} else {
			set[current.Next.Val] = true
			current = current.Next
		}
	}

	return head
}
