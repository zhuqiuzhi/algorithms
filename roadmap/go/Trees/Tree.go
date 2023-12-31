package roadmap

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func size(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return size(root.Left) + size(root.Right) + 1
}
