package roadmap

/*
Leetcode Problem: https://leetcode.com/problems/diameter-of-binary-tree/
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.
*/
// 难点: 最大路径不一定要经过根结点, 所以要计算所有结点的的最大路径
func diameterOfBinaryTree(root *TreeNode) int {
	var ret int

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMaxDepth := maxDepth(root.Left)
		rightMaxDepth := maxDepth(root.Right)

		// update result
		// +1 是因为把根结点也计算进去, -1 是因为路径长度等于结点数-1
		ret = max(ret, leftMaxDepth+rightMaxDepth+1-1)

		return max(leftMaxDepth, rightMaxDepth) + 1
	}

	maxDepth(root)
	return ret
}
