package roadmap

func isValid(s string) bool {
	pairs := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if _, ok := pairs[char]; ok {
			// left, push it onto stack
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		// pop
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pairs[top] != char {
			// not bracket or not valid
			return false
		}
	}

	return len(stack) == 0
}
