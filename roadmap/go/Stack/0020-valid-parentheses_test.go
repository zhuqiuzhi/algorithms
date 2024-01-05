package roadmap

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"single", "()", true},
		{"mutil", "()[]{}", true},
		{"[}", "[}", false},
		{"right empty", "(", false},
		{"left empty", ")", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.arg); got != tt.want {
				t.Errorf("case %s: isValid('%s') = %v, want %v", tt.name, tt.arg, got, tt.want)
			}
		})
	}
}
