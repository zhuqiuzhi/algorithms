package roadmap

// Leetcode Problem: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree
// and inorder is the inorder traversal of the same tree,
// construct and return the binary tree.
// Constraints:
// 1 <= preorder.length <= 3000
// inorder.length == preorder.length
// -3000 <= preorder[i], inorder[i] <= 3000
// preorder and inorder consist of unique values.
// 难点: 前序遍历序列只能判断出根结点和左子树的根结点, 右子树的根结点怎么判断
// Input: preorder = [3,9,11,20,15,7], inorder = [11,9,3,15,20,7]
// 根结点是 3, 在中序遍历中找到, 左子树中序遍历是 [11,9], 右子树中序遍历是 [15,20,7]
// 根据左子树的长度 2, 判断出左子树的前序遍历是[9, 11], 右子树的前序序列是 [20,15, 7]
func buildTree(preorder []int, inorder []int) *TreeNode {
	// base case
	if len(preorder) == 0 {
		return nil
	}

	// pop root
	rootValue := preorder[0]
	preorder = preorder[1:]
	root := &TreeNode{Val: rootValue}

	// python 可以直接写成: index = inorder.index(rootValue)
	var index int
	for index = 0; index < len(inorder); index++ {
		if inorder[index] == rootValue {
			break
		}
	}

	// 关键技巧: 根据中序序列判断出左子树和右子树, 然后根据左子树的长度, 得到左子树的前序遍历序列和中序遍历序列
	leftSize := len(inorder[:index])
	root.Left = buildTree(preorder[:leftSize], inorder[:index])
	root.Right = buildTree(preorder[leftSize:], inorder[index+1:])

	return root
}
