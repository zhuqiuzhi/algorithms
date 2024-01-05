package roadmap

// Given a string s, find the length of the longest
// substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	var charSet = make(map[rune]struct{}) // 窗口内元素集合
	var left, res = 0, 0
	inFunc := func(r rune) bool {
		_, ok := charSet[r]
		return ok
	}
	runes := []rune(s)

	for i, r := range runes {
		// 特殊技巧: 通过判断重复字符 r 是否还在元素集合内来缩小窗口
		for inFunc(r) {
			delete(charSet, runes[left])
			left++
		}
		charSet[r] = struct{}{}
		res = max(res, i-left+1)
	}

	return res
}
