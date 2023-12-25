package roadmap

// Leetcode Problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var minPrice = prices[0]
	var maxProfit int
	for _, price := range prices {
		// 循环不变式: minPrice 是当前价格之前的最低价格
		if price < minPrice {
			// 不可能产生最大利润
			minPrice = price
		} else {
			// 可能产生最大利润
			maxProfit = max(maxProfit, price-minPrice)
		}
	}
	return maxProfit
}
