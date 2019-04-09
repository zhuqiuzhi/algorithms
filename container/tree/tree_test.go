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
	node3 := &Node{
		Val: 3,
	}
	node2 := &Node{
		Val:  2,
		Left: node3,
	}
	root := &Node{
		Val:   1,
		Right: node2,
	}
	tcs := []struct {
		testFunc func(root *Node) []int
		root     *Node
		want     []int
	}{
		{preorderTraversal, root, []int{1, 2, 3}},
		{postorderTraversal, root, []int{3, 2, 1}},
		{inorderTraversal, root, []int{1, 3, 2}},
	}

	for _, tc := range tcs {
		got := tc.testFunc(tc.root)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s got %v want %v", GetFunctionName(tc.testFunc), got, tc.want)
		}
	}
}
