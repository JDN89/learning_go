package grind

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case until root is nil or equals to p or q
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// left and right or not nil and have same ancestor
	if left != nil && right != nil {
		return root
	}

	//      root
	//     /   \
	//    left  nil

	if left != nil {
		return left
	}
	//      root
	//     /   \
	//    nil  right

	return right
}
