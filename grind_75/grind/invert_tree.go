package grind

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// queue is BFS
func invertTree(root *TreeNode) *TreeNode {
	savedRoot := root
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	// queue[0] = root

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left == nil && current.Right == nil {
			continue
		}
		current.Left, current.Right = current.Right, current.Left
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}

		continue

	}
	return savedRoot
}
