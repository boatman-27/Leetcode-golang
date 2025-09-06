package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeSum struct {
	node       *TreeNode
	currentSum int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	q := []*nodeSum{{root, root.Val}}
	for len(q) != 0 {
		current := q[0]
		q = q[1:]

		if current.node.Left == nil && current.node.Right == nil && current.currentSum == targetSum {
			return true
		}

		if current.node.Left != nil {
			q = append(q, &nodeSum{current.node.Left, current.currentSum + current.node.Left.Val})
		}
		if current.node.Right != nil {
			q = append(q, &nodeSum{current.node.Right, current.currentSum + current.node.Right.Val})
		}
	}

	return false
}
