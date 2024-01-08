package roadmap

// Leetcode Problem: https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
func numberOfWays(startPos int, endPos int, k int) int {
	// n 是从 startPos 到 endPos 的步数
	var dfs func(startPos, endPos, n int) int
	memo := make(map[int]map[int]int)

	dfs = func(startPos, endPos, n int) int {
		if m, ok := memo[startPos]; ok {
			if ret, ok := m[n]; ok {
				return ret
			}
		}

		if startPos == endPos {
			if n == 0 || n%2 == 0 {
				return 1
			} else {
				return 0
			}
		}

		ret := dfs(startPos-1, endPos, n-1) + dfs(startPos+1, endPos, n-1)
		_, ok := memo[startPos]
		if !ok {
			memo[startPos] = make(map[int]int)
		}
		memo[startPos][n] = ret
		return ret
	}

	return dfs(startPos, endPos, k)
}

// 时间复杂度太高, 导致会超时
func numberOfWaysDfs(startPos int, endPos int, k int) int {
	var ret int

	var dfs func(startPos, endPos, i int)
	dfs = func(startPos, endPos, i int) {
		if startPos == endPos && i == k {
			ret++
			return
		}
		if i > k {
			return
		}
		// left
		dfs(startPos-1, endPos, i+1)
		// right
		dfs(startPos+1, endPos, i+1)
	}
	dfs(startPos, endPos, 0)

	return ret
}
