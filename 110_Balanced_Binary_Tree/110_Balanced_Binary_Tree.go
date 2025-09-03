package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
