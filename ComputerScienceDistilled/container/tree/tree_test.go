package tree

import (
	"reflect"
	"runtime"
	"testing"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestTraversal(t *testing.T) {
	node3 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val:  2,
		Left: node3,
	}
	root := &TreeNode{
		Val:   1,
		Right: node2,
	}
	tcs := []struct {
		testFunc func(root *TreeNode) []int
		root     *TreeNode
		want     []int
	}{
		{preorderTraversalR, root, []int{1, 2, 3}},
		{preorderTraversal, root, []int{1, 2, 3}},

		{postorderTraversalR, root, []int{3, 2, 1}},
		{postorderTraversal, root, []int{3, 2, 1}},

		{inorderTraversalR, root, []int{1, 3, 2}},
		{inorderTraversal, root, []int{1, 3, 2}},
		{inorderTraversalNeetcode, root, []int{1, 3, 2}},
	}

	for _, tc := range tcs {
		got := tc.testFunc(tc.root)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s got %v want %v", GetFunctionName(tc.testFunc), got, tc.want)
		}
	}
}
