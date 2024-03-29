package roadmap

import (
	"unicode"
)

func isPalindrome(s string) bool {
	// 重要语言技巧: 将字符串转为 []rune
	runes := []rune(s)

	i, j := 0, len(runes)-1
	for i < j {
		start := runes[i]
		if !(unicode.IsDigit(start) || unicode.IsLetter(start)) {
			i++
			continue
		}

		end := runes[j]
		if !(unicode.IsDigit(end) || unicode.IsLetter(end)) {
			j--
			continue
		}

		if unicode.ToLower(start) != unicode.ToLower(end) {
			return false
		}

		i++
		j--
	}
	return true
}
