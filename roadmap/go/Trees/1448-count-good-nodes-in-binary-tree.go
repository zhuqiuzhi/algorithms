package roadmap

// Leetcode Problem: https://leetcode.com/problems/count-good-nodes-in-binary-tree/
// Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.
// Constraints:
// The number of nodes in the binary tree is in the range [1, 10^5].
// Each node's value is between [-10^4, 10^4].
// 难点: 将每个结点需要和从根结点到该结点上路径的所有结点比较, 转换为, 只需要和路径上最大值比较
// 微软面试考察最多的一道题
func goodNodes(root *TreeNode) int {
	var ret int

	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			max = root.Val
			ret++
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}

	// 易错点, 结点值可能是负数, max 不能传 0
	dfs(root, root.Val)

	return ret
}
