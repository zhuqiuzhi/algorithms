package roadmap

// Leetcode Problem: https://leetcode.com/problems/binary-tree-right-side-view/
// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.
// 考虑: [1,2], [1,2,3,4]
// 技巧: 按行遍历树, 然后将树每一行最后一个结点就是从右侧看到的第一个结点
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		curLen := len(queue)
		// 难点: 这里需要先将 result 数组长度增加 1, 末尾增加一个空 slice
		// 技巧: 每次将一行结点添加完, 再遍历下一行结点
		for i := 0; i < curLen; i++ {
			cur := queue[i]
			if cur == nil {
				continue
			}
			// 将每行最后一个结点值加到 result 中
			if i == curLen-1 {
				result = append(result, cur.Val)
			}

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
