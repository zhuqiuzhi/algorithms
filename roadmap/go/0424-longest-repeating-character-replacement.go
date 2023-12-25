package roadmap

// Letcode Problem: https://leetcode.com/problems/longest-repeating-character-replacement/
// 难点是如何处理明明有机会用,但 k 次机会没有用完的情况, 例如 s = "ABBB",k = 2
// 以下是错误的实现:
func characterErrorReplacement(s string, k int) int {
	if s == "" {
		return 0
	}

	runes := []rune(s)
	var res, left, right int
	replaceNum := k
	var c rune
	for i, r := range runes {
		if i == 0 {
			c = r
			left = 0
			right = 0
		} else if r == c {
			right++
		} else {
			if replaceNum != 0 {
				right++
				replaceNum--
			} else {
				left = i
				right = i
				c = r
				replaceNum = k
			}
		}

		res = max(res, right-left+1)
	}

	return res
}

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res := 0

	left := 0
	maxf := 0

	for right := range s {
		// update count map
		count[s[right]] = 1 + count[s[right]]

		// update max frequency that char in windows
		maxf = max(maxf, count[s[right]])

		// 判断窗口内是否可以通过至多 k 次替换将其他字符(非出现最大次数的字符)替换掉
		if (right-left+1)-maxf > k {
			// 将窗口最左边的字符出现频率减1,并将窗口左端向右移动一位,维持窗口内是有效的
			count[s[left]] -= 1
			// 没有必要减小 maxf, 因为 maxf 减少不会影响最终的结果
			// maxf = getMaxf(count)
			left++
		}

		res = max(res, right-left+1)
	}
	return res
}

func getMaxf(count map[byte]int) int {
	var maxf int
	for _, count := range count {
		if count > maxf {
			maxf = count
		}
	}
	return maxf
}
