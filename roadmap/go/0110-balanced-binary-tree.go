package roadmap

// Leetcode Problem: https://leetcode.com/problems/balanced-binary-tree/
// Given a binary tree, determine if it is height-balanced.
// For this problem, the height-balanced tree is defined as:
// a binary tree in which the left and right subtrees of every node differ
// in height by no more than 1.
// 注意空树是高度平衡树, 且需要每个结点都是高度平衡, 不能只有根结点
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var maxDepth func(root *TreeNode) int
	var notBalancedNode int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMaxDepth := maxDepth(root.Left)
		rightMaxDepth := maxDepth(root.Right)
		if leftMaxDepth > rightMaxDepth && leftMaxDepth-rightMaxDepth > 1 {
			notBalancedNode++
		}
		if rightMaxDepth > leftMaxDepth && rightMaxDepth-leftMaxDepth > 1 {
			notBalancedNode++
		}
		return max(leftMaxDepth, rightMaxDepth) + 1
	}

	maxDepth(root)
	return notBalancedNode == 0
}
