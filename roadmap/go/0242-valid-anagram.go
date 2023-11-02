package roadmap

func isAnagramWithMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var countS, countT = map[rune]int{}, map[rune]int{}
	for _, r := range s {
		countS[r]++
	}
	for _, r := range t {
		countT[r]++
	}
	for r, count := range countS {
		if countT[r] != count {
			return false
		}
	}
	return true
}

func isAnagramWithSlice(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 注意不需要使用两个数组
	var freqs [26]int
	//  一次循环遍历两个字符串, 相互抵消
	for i := 0; i < len(s); i++ {
		freqs[s[i]-'a']++
		freqs[t[i]-'a']--
	}

	for _, freq := range freqs {
		if freq != 0 {
			return false
		}
	}

	return true
}
