package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeDepth struct {
	node  *TreeNode
	depth int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []NodeDepth{{root, 1}}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if current.node.Left == nil && current.node.Right == nil {
			return current.depth
		}
		if current.node.Left != nil {
			q = append(q, NodeDepth{current.node.Left, current.depth + 1})
		}
		if current.node.Right != nil {
			q = append(q, NodeDepth{current.node.Right, current.depth + 1})
		}
	}

	return 0
}
