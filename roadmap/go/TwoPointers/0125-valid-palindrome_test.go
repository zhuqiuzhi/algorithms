package roadmap

import "testing"

func TestIsPalindrome(t *testing.T) {
	tcs := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}
	for _, tc := range tcs {
		got := isPalindrome(tc.in)
		if got != tc.want {
			t.Errorf("isPalindrome(%s) got %t, want %t", tc.in, got, tc.want)
		}
	}
}
