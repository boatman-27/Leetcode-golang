package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	return 1 + max(leftHeight, rightHeight)
}
