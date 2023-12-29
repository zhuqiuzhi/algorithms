package roadmap

// Leetcode Problem: https://leetcode.com/problems/binary-tree-level-order-traversal/
// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		curLen := len(queue)
		// 难点: 这里需要先将 result 数组长度增加 1, 末尾增加一个空 slice
		result = append(result, []int{})
		// 技巧: 每次将一行结点添加完, 再遍历下一行结点
		for i := 0; i < curLen; i++ {
			cur := queue[i]
			if cur == nil {
				continue
			}

			result[len(result)-1] = append(result[len(result)-1], cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 易错点: 可能左右子树都为空, 此时 queue 队列长度没有增大
		if len(queue) > curLen {
			queue = queue[curLen:]
		} else {
			queue = queue[:0]
		}
	}

	return result
}
