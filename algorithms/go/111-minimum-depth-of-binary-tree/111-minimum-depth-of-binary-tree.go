package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	left_depth := minDepth(root.Left)
	right_depth := minDepth(root.Right)
	if left_depth > 0 && right_depth > 0 {
		depth = 1 + Min(left_depth, right_depth)
	} else {
		depth = 1 + left_depth + right_depth
	}
	return depth
}

// Max returns the larger of x or y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
