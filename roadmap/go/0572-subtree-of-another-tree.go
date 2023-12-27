package roadmap

// Leetcode Problem: https://leetcode.com/problems/subtree-of-another-tree/
// Given the roots of two binary trees root and subRoot,
// return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants.
// The tree tree could also be considered as a subtree of itself.
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val {
		isSame := isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
		if isSame {
			return true
		}
		// 容易出错点, 即使根结点的值相同, 也有可能是和左子树或者右子树中的子树相同
		// 例如: root=[1,1], subRoot=[1]
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}
