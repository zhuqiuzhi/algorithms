from typing import List


class Solution:
    def maxProfitBrute(self, prices: List[int]) -> int:
        """ Best Time to Buy and Sell Stock 

        >>> Solution().maxProfitBrute([7,1,5,3,6,4])
        5
        >>> Solution().maxProfitBrute([7,6,4,3,1])
        0
        """
        maxProfit = 0
        for start in range(0, len(prices)):
            for end in range(start+1, len(prices)):
                diff = prices[end] - prices[start]
                if diff > maxProfit:
                    maxProfit = diff

        return maxProfit

    def maxProfit(self, prices: List[int]) -> int:
        """ Best Time to Buy and Sell Stock

        >>> Solution().maxProfit([7,1,5,3,6,4])
        5
        >>> Solution().maxProfit([7,6,4,3,1])
        0
        """
        result = 0
        lowest = prices[0]
        for price in prices:
            if price < lowest:
                lowest = price
            else:
                result = max(result, price - lowest)

        return result
