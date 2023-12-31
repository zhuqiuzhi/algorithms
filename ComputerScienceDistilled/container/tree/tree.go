package tree

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先根遍历的递归版本
func preorderTraversalRecursion(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}

	values = append(values, root.Val)
	// 易错点: 没有将返回的 values 保存
	values = preorderTraversalRecursion(root.Left, values)
	values = preorderTraversalRecursion(root.Right, values)

	return values
}

func preorderTraversalR(root *TreeNode) []int {
	var ret []int
	ret = preorderTraversalRecursion(root, ret)
	return ret
}

// 先根遍历的迭代版本
// preorderTraversal 使用先根遍历的方法遍历二叉树, 返回遍历的序列
func preorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack = []*TreeNode{root}
	var n *TreeNode

	for len(stack) != 0 {
		n, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if n == nil {
			continue
		}
		ret = append(ret, n.Val)
		// 注意这里是先放右子树, 然后再放左子树
		// 因为栈的特点是先进后出, 所以会先处理完左子树的所有结点再处理右子树的
		stack = append(stack, n.Right)
		stack = append(stack, n.Left)
	}

	return ret
}

// 后根遍历的递归版本
func postorderTraversalRecursion(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}

	// 易错点: 没有将返回的 values 保存
	values = postorderTraversalRecursion(root.Left, values)
	values = postorderTraversalRecursion(root.Right, values)
	values = append(values, root.Val)

	return values
}

func postorderTraversalR(root *TreeNode) []int {
	var ret []int
	ret = postorderTraversalRecursion(root, ret)
	return ret
}

const (
	visit = 0
	print = 1
)

// Guide 用保存后根遍历和中根遍历时对结点进行处理的操作
// 因为后根遍历和中根遍历在访问完根结点时,还没有处理完, 还需要放回栈中之后再处理(打印)
// 所以还需要区分是访问结点还是打印结点
type Guide struct {
	op   int // 0: visit 1: print
	node *TreeNode
}

// 后根遍历的迭代版本
// postorderTraversal 使用后根遍历的方法遍历二叉树, 返回遍历的序列
func postorderTraversal(root *TreeNode) []int {
	var ret []int
	var stack = []Guide{{op: visit, node: root}}
	var current Guide

	for len(stack) != 0 {
		// pop
		current, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if current.node == nil {
			continue
		}
		if current.op == print {
			// print
			ret = append(ret, current.node.Val)
		} else {
			// visit
			stack = append(stack, Guide{print, current.node})
			stack = append(stack, Guide{visit, current.node.Right})
			stack = append(stack, Guide{visit, current.node.Left})
		}
	}

	return ret
}

// 中根遍历的递归版本
func inorderTraversalRecursion(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}

	// 易错点: 没有将返回的 values 保存
	values = inorderTraversalRecursion(root.Left, values)
	values = append(values, root.Val)
	values = inorderTraversalRecursion(root.Right, values)

	return values
}

func inorderTraversalR(root *TreeNode) []int {
	var ret []int
	ret = inorderTraversalRecursion(root, ret)
	return ret
}

// 中根遍历的迭代版本
// inorderTraversal 使用中根遍历的方法遍历二叉树, 返回遍历的序列
func inorderTraversal(root *TreeNode) []int {
	var ret []int
	// 将根结点压入栈中
	var stack = []Guide{{op: visit, node: root}}
	var current Guide

	for len(stack) != 0 {
		// pop
		current, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if current.node == nil {
			continue
		}
		if current.op == 1 {
			// print
			ret = append(ret, current.node.Val)
		} else {
			// visit
			stack = append(stack, Guide{visit, current.node.Right})
			stack = append(stack, Guide{print, current.node})
			stack = append(stack, Guide{visit, current.node.Left})
		}
	}

	return ret
}

func inorderTraversalNeetcode(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode

	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		ret = append(ret, cur.Val)

		cur = cur.Right
	}

	return ret
}
