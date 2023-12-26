package roadmap

// Leetcode Problem: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
