package LRU

import (
	"reflect"
	"testing"
)

func TestLRU(t *testing.T) {
	tcs := []struct {
		cap   int
		ops   []string
		input [][]int
		want  []int
	}{
		{
			cap:   2,
			ops:   []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
			input: [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			want:  []int{0, 0, 1, 0, -1, 0, -1, 3, 4},
		},
		{
			cap:   2,
			ops:   []string{"put", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			input: [][]int{{1, 1}, {2, 2}, {1, 5}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			want:  []int{0, 0, 0, 5, 0, -1, 0, -1, 3, 4},
		},
	}

	for i, tc := range tcs {
		if len(tc.ops) != len(tc.input) || len(tc.ops) != len(tc.want) {
			t.Fatalf("test case %d' len(ops)=%d, len(input)=%d, len(want)=%d", i, len(tc.ops), len(tc.input), len(tc.want))
		}

		var got []int
		lru := New(tc.cap)
		for j, op := range tc.ops {
			switch op {
			case "put":
				lru.Put(tc.input[j][0], tc.input[j][1])
				got = append(got, 0)
			case "get":
				got = append(got, lru.Get(tc.input[j][0]))
			}
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("case %d got %v", i, got)
		}
	}
}
