package roadmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var InitGlobal = 10
var NotInitGlobal int

func TestCharacterReplacement(t *testing.T) {
	type testCase struct {
		Name string

		S string
		K int

		ExpectedInt int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualInt := characterReplacement(tc.S, tc.K)

			assert.Equal(t, tc.ExpectedInt, actualInt)
		})
	}

	validate(t, &testCase{
		Name:        "ABAB",
		S:           "ABAB",
		K:           2,
		ExpectedInt: 4,
	})

	validate(t, &testCase{
		Name:        "AABABBA",
		S:           "AABABBA",
		K:           1,
		ExpectedInt: 4,
	})
	validate(t, &testCase{
		Name:        "ABBB",
		S:           "ABBB",
		K:           2,
		ExpectedInt: 4,
	})

	validate(t, &testCase{
		Name:        "ABBBCDE",
		S:           "ABBBCDE",
		K:           1,
		ExpectedInt: 4,
	})
	validate(t, &testCase{
		Name:        "ABBBEEEE",
		S:           "ABBBEEEE",
		K:           1,
		ExpectedInt: 5,
	})

}
