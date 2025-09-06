package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	dummy := &ListNode{0, head}
	current := dummy

	for current.Next != nil {
		if set[current.Next.Val] {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}
