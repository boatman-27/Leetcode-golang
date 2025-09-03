package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToTail(val int, currentResult *ListNode) {
	newNode := &ListNode{Val: val}
	if currentResult == nil {
		currentResult = newNode
		return
	}

	current := currentResult
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			AddToTail(list1.Val, result)
			list1 = list1.Next
		} else {
			AddToTail(list2.Val, result)
			list2 = list2.Next
		}
	}

	// Add leftovers
	for list1 != nil {
		AddToTail(list1.Val, result)
		list1 = list1.Next
	}

	for list2 != nil {
		AddToTail(list2.Val, result)
		list2 = list2.Next
	}
	return result
}
