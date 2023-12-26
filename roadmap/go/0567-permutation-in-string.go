package roadmap

// Leetcode Problem: https://leetcode.com/problems/permutation-in-string/
// 关键点, 滑动窗口大小应该是固定的, 并且 s1 可能有重复的字符
// 难点: 滑动窗口后, 怎么更新 map, 怎么不更新 map 直接计算匹配数
// 技巧: 根据窗口右移前后对比, 来更新匹配数
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var countS1, countS2 [26]int // 字母-'a' => 字母出现次数
	for i := 0; i < len(s1); i++ {
		countS1[s1[i]-'a'] += 1
		countS2[s2[i]-'a'] += 1
	}

	var match int // countS1 和 count S2 字母出现次数的匹配个数
	// 初始化匹配数: 第一个窗口的匹配数
	for i := 0; i < len(countS1); i++ {
		if countS1[i] == countS2[i] {
			match += 1
		}
	}

	left := 0

	// countS1 保持不变, 固定窗口大小移动时, 更新 countS2, 然后根据窗口右移前后对比, 更新 match 数
	for right := len(s1); right < len(s2); right++ {
		if match == 26 {
			return true
		}

		index := s2[right] - 'a'
		countS2[index] += 1
		if countS1[index] == countS2[index] {
			// 说明窗口右移后, 和前一个窗口相比, 匹配数增加了 1(因为加1后才想等, 说明之前这个字母是不匹配)
			match += 1
		} else if countS1[index] == countS2[index]-1 {
			// 说明窗口右移前, 这个字母是匹配的(现在减去1等于s1中这个字母出现的次数)
			match -= 1
		}

		// 去掉窗口最左端, match 需要变化
		index = s2[left] - 'a'
		countS2[index] -= 1
		if countS1[index] == countS2[index]+1 {
			// 说明原来是匹配的, 现在窗口去掉了最左端, 所以匹配数需要减 1
			match -= 1
		} else if countS1[index] == countS2[index] {
			// 说明窗口右移后, 这个字母是匹配的, 原来是不匹配的, 所以匹配数需要增加 1
			match++
		}

		left += 1
	}

	return match == 26
}
