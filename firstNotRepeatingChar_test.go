package algorithms

import "testing"

func TestFirstNotRepeatingChar(t *testing.T) {
	tcs := []struct {
		in   string
		want byte
	}{
		{"abefafbecc", 0},       /* 没有不重复的字符 */
		{"abfacfb", byte('c')},  /* 只有一个不重复的字符 */
		{"abfacfbd", byte('c')}, /* 有两个不重复的字符 */
	}

	for _, tc := range tcs {
		got := FirstNotRepeatingChar(tc.in)
		if got != tc.want {
			t.Errorf("FirstNotRepeatingChar(%s) got %d want %d", tc.in, got, tc.want)
		}
	}
}
