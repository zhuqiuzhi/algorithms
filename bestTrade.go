// BEST TRADE: You have the daily prices of gold for a interval of time.
// You want to find two days in this interval such that if you had bought then sold gold at
// those dates, you’d have made the maximum possible profit.

package algorithms

// trade 使用分治法自顶向下的方式解决 "BEST TRADE" 问题
func trade(prices []int) []int {
	// TODO
	return []int{}
}

// trade_dp 使用动态规划法自底向上的方式解决 "BEST TRADE" 问题
// P(n) is the price on the nth day.
// B(n) is the best day to buy if we’re selling on the nth day.
// B(n) = n if P(n) < P(B(n-1))
// B(n) = B(n-1) if P(n) >= P(B(n-1))
func tradeDP(prices []int) []int {
	// b 记录每日卖出时的最佳买入时间
	var b = make(map[int]int)
	b[0] = 0

	sellDay, bestProfit := 0, 0
	for i := 1; i < len(prices); i++ {
		// 根据第 i 天的价格和前一天卖出时的最佳买入价格比较，得到第 i 天最佳买入的时间
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
// 只需要存储相对于目前最佳卖出日的最佳买入日即可
// 最佳交易组合出现在 [0, i-1] 中的最佳交易组合和 (b, i) 交易组合中
// b 是比 [0, i-1] 中的交易组合的买入日价格更低的日期中最晚的日期
func tradeKadane(prices []int) []int {
	buyDay, sellDay, bestProfit := 0, 0, 0

	// b 记录比 buyDay 日价格更低的日期
	b := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[buyDay] {
			b = i
		}
		profit := prices[i] - prices[b]
		if profit > bestProfit {
			buyDay = b
			sellDay = i
			bestProfit = profit
		}
	}
	return []int{sellDay, buyDay}
}
