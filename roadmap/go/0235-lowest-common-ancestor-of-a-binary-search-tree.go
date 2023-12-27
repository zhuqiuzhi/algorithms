package roadmap

// Leetcode Problem: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
// According to the definition of LCA on Wikipedia:
// “The lowest common ancestor is defined between two nodes p
// and q as the lowest node in T that has both p and q as
// descendants (where we allow a node to be a descendant of itself).”
// Constraints:
// All Node.val are unique.
// p != q
// p and q will exist in the BST.
// 注意 root 是二叉搜索树且 p 和 q 总是在树里面
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root

	// 易错点, 这里没有是要和 cur.Val 比较而不是 root.Val 比较
	for cur != nil {
		if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}

	return nil
}
