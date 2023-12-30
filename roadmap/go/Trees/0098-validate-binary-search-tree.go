package roadmap

import "math"

// Leetcode Problem: https://leetcode.com/problems/validate-binary-search-tree/
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, min int, max int) bool
	dfs = func(root *TreeNode, min int, max int) bool {
		if root == nil {
			return true
		}
		// 易错点: 等于最小值或者最大值也不行
		if root.Val <= min || root.Val >= max {
			return false
		}

		// 难点: 左结点也需要大于根结点的要求的最喜
		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
