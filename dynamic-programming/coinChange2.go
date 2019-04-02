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

	// Input:
	// amount=5, coins=[]int{1,2,5}
	// Output:
	// coin: 1, ways: [1 1 1 1 1 1]
	// coin: 2, ways: [1 1 2 2 3 3]
	// coin: 5, ways: [1 1 2 2 3 4]

	return store[amount]
}
