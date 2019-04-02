// Coin Change 2
// You are given coins of different denominations and a total amount of money.
// Write a function to compute the number of combinations that make up that amount.
// You may assume that you have infinite number of each kind of coin.
package dynamic

import (
	"sort"
)

// change2 通过动态规划解决硬币兑换的组合个数问题
// https://www.youtube.com/watch?v=_fgjrs570YE
func change2(amount int, coins []int) int {
	sort.Ints(coins)

	store := make([]int, amount+1)

	// base case
	store[0] = 1

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i >= coin {
				store[i] += store[i-coin]
			}
		}
		//fmt.Printf("coin: %d, ways: %v\n", coin, store)
	}

	// 解题思路
	// 使用 {..., coin} 兑换 amount:
	// 兑换 amount 的组合数分为两种情况，一种使用 coin，一种没有使用 coin
	// 使用 coin 时兑换的组合数:
	//    如果 amount 大于 coin，则表示使用 coin 时的组合数等于不使用该coin时兑换 amount-coin 时的组合数;
	//    因为这些组合只要再加上 coin，就能得到 amount;
	//    amount 等于 coin 时，则表示使用coin时的，组合数就是 1;
	//    否则使用 coin 时的组合数为 0;
	// 不使用 coin 时兑换 amount 的组合数就是在上一次迭代中计算出的 store[i]
	//
	// 总的组合数= 使用coin时兑换的组合数 + 不使用 coin 时兑换 amount 的组合数
	//
	// Input:
	// amount=5, coins=[]int{1,2,5}
	// Output:         0 1 2 3 4 5
	// coin: 1, ways: [1 1 1 1 1 1]
	// coin: 2, ways: [1 1 2 2 3 3]
	// coin: 5, ways: [1 1 2 2 3 4]

	return store[amount]
}
