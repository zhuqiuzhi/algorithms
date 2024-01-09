package roadmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfWaysDfs(t *testing.T) {
	type testCase struct {
		Name string

		StartPos int
		EndPos   int
		K        int

		ExpectedInt int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualInt := numberOfWays(tc.StartPos, tc.EndPos, tc.K)

			assert.Equal(t, tc.ExpectedInt, actualInt)
		})
	}

	validate(t, &testCase{StartPos: 1, EndPos: 2, K: 3, ExpectedInt: 3})
	validate(t, &testCase{StartPos: 2, EndPos: 5, K: 10, ExpectedInt: 0})
	validate(t, &testCase{StartPos: 1, EndPos: 1000, K: 999, ExpectedInt: 1})
}
