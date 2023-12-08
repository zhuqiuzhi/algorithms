// BEST TRADE: You have the daily prices of gold for a interval of time.
// You want to find two days in this interval such that if you had bought then sold gold at
// those dates, you’d have made the maximum possible profit.

package dynamic

// trade 使用分治法自顶向下的方式解决 "BEST TRADE" 问题
func trade(prices []int) []int {
	// TODO
	return []int{}
}

// trade_dp 使用动态规划法自底向上的方式解决 "BEST TRADE" 问题
func tradeDP(prices []int) []int {
	// b[i] 记录 [0,i] 中价格最低的日期的下标
	var b = make(map[int]int)
	b[0] = 0

	sellDay, bestProfit := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[b[i-1]] {
			b[i] = i
		} else {
			b[i] = b[i-1]
		}

		profit := prices[i] - prices[b[i]]
		if profit > bestProfit {
			bestProfit = profit
			sellDay = i
		}
	}
	return []int{sellDay, b[sellDay]}
}

// trade_kadane 使用1984 年 Jay Kadane 教授提出的算法，其时间复杂度是 O(n)，但空间复杂度只要 O(1)
// 也是《剑指Offer》第二版第 63 题的解法
func tradeKadane(prices []int) []int {
	buyDay, sellDay, bestProfit := 0, 0, 0

	min := 0
	for i := 1; i < len(prices); i++ {
		// 循环不变式: min 是 [0, i] 中价格最小的下标
		if prices[i] < prices[min] {
			min = i
		}

		profit := prices[i] - prices[min]
		if profit > bestProfit {
			sellDay = i
			buyDay = min
			bestProfit = profit
		}
	}
	return []int{sellDay, buyDay}
}
