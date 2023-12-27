package tree

// Node is a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// 先根遍历的递归版本
func preorderTraversalRecursion(root *Node, values []int) []int {
	if root == nil {
		return values
	}

	values = append(values, root.Val)
	// 易错点: 没有将返回的 values 保存
	values = preorderTraversalRecursion(root.Left, values)
	values = preorderTraversalRecursion(root.Right, values)

	return values
}

func preorderTraversalR(root *Node) []int {
	var ret []int
	ret = preorderTraversalRecursion(root, ret)
	return ret
}

// 先根遍历的迭代版本
// preorderTraversal 使用先根遍历的方法遍历二叉树, 返回遍历的序列
func preorderTraversal(root *Node) []int {
	var ret []int
	var stack = []*Node{root}
	var n *Node

	for len(stack) != 0 {
		n, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if n == nil {
			continue
		}
		ret = append(ret, n.Val)
		stack = append(stack, n.Right)
		stack = append(stack, n.Left)
	}

	return ret
}

// 后根遍历的递归版本
func postorderTraversalRecursion(root *Node, values []int) []int {
	if root == nil {
		return values
	}

	// 易错点: 没有将返回的 values 保存
	values = postorderTraversalRecursion(root.Left, values)
	values = postorderTraversalRecursion(root.Right, values)
	values = append(values, root.Val)

	return values
}

func postorderTraversalR(root *Node) []int {
	var ret []int
	ret = postorderTraversalRecursion(root, ret)
	return ret
}

// Guide 用保存后根遍历和中根遍历时对结点进行处理的操作
type Guide struct {
	op   int // 0: visit 1: print
	node *Node
}

// 后根遍历的迭代版本
// postorderTraversal 使用后根遍历的方法遍历二叉树, 返回遍历的序列
func postorderTraversal(root *Node) []int {
	var ret []int
	var stack = []Guide{{op: 0, node: root}}
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
			stack = append(stack, Guide{1, current.node})
			stack = append(stack, Guide{0, current.node.Right})
			stack = append(stack, Guide{0, current.node.Left})
		}
	}

	return ret
}

// 中根遍历的递归版本
func inorderTraversalRecursion(root *Node, values []int) []int {
	if root == nil {
		return values
	}

	// 易错点: 没有将返回的 values 保存
	values = inorderTraversalRecursion(root.Left, values)
	values = append(values, root.Val)
	values = inorderTraversalRecursion(root.Right, values)

	return values
}

func inorderTraversalR(root *Node) []int {
	var ret []int
	ret = inorderTraversalRecursion(root, ret)
	return ret
}

// 中根遍历的迭代版本
// inorderTraversal 使用中根遍历的方法遍历二叉树, 返回遍历的序列
func inorderTraversal(root *Node) []int {
	var ret []int
	// 将根结点压入栈中
	var stack = []Guide{{op: 0, node: root}}
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
			stack = append(stack, Guide{0, current.node.Right})
			stack = append(stack, Guide{1, current.node})
			stack = append(stack, Guide{0, current.node.Left})
		}
	}

	return ret
}
