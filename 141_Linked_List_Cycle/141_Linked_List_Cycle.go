package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	set := map[*ListNode]bool{}
	current := head

	for current != nil {
		if _, ok := set[current]; ok {
			return true
		}
		set[current] = true
		current = current.Next
	}
	return false
}
