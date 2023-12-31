package roadmap

// Leetcode Problem: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Given the root of a binary search tree, and an integer k,
// return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
// 递归版解法
func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode, k int) *TreeNode
	dfs = func(root *TreeNode, k int) *TreeNode {
		if root == nil {
			return nil
		}

		t := size(root.Left)
		if t > k {
			// 左子树中结点个数大于 k 个, 第 k 个(0-indexed) 是在左子树中
			return dfs(root.Left, k)
		} else if t < k {
			// 递归地在右子树中寻找左子树中结点数是 k-1-t 的结点, 因为当前根结点加上左子数中有 t+1 个结点
			return dfs(root.Right, k-1-t)
		} else {
			// 左子树中结点个数等于 k 个, 第 k 个(0-indexed) 就是根结点
			return root
		}
	}

	node := dfs(root, k-1)
	if node != nil {
		return node.Val
	}

	return -1
}

// 深入优先搜索
func kthSmallestR(root *TreeNode, k int) int {
	var n, ret int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		n++
		if n == k {
			ret = root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return ret
}

type Guide struct {
	P       *TreeNode
	IsVisit bool
}

// 万能迭代版解法
func kthSmallestI(root *TreeNode, k int) int {
	var num int

	var stack = []*Guide{&Guide{P: root, IsVisit: true}}
	var n *Guide
	for len(stack) != 0 {
		// pop stack
		n, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if !n.IsVisit {
			num++
			if num == k {
				return n.P.Val
			}
			continue
		}

		// 注意这里是先放右子树, 然后再放左子树
		// 因为栈的特点是先进后出, 所以会先处理完左子树的所有结点再处理右子树的
		if n.P.Right != nil {
			stack = append(stack, &Guide{P: n.P.Right, IsVisit: true})
		}
		stack = append(stack, &Guide{P: n.P, IsVisit: false})
		if n.P.Left != nil {
			stack = append(stack, &Guide{P: n.P.Left, IsVisit: true})
		}
	}

	return 0
}

func kthSmallestNeetcode(root *TreeNode, k int) int {
	var n int
	var stack []*TreeNode
	cur := root

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			// push into stack
			stack = append(stack, cur)
			cur = cur.Left
		}

		// pop stack
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n++
		if n == k {
			return cur.Val
		}

		cur = cur.Right
	}

	return 0
}
