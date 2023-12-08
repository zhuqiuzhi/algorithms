package tree

// Node is a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

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

type Guide struct {
	op   int // 0: visit 1: print
	node *Node
}

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

// inorderTraversal 使用中根遍历的方法遍历二叉树, 返回遍历的序列
func inorderTraversal(root *Node) []int {
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
			stack = append(stack, Guide{0, current.node.Right})
			stack = append(stack, Guide{1, current.node})
			stack = append(stack, Guide{0, current.node.Left})
		}
	}

	return ret
}
