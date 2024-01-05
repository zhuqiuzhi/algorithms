from typing import List


class Solution:
    """
    >>> s = Solution()
    >>> s.twoSum([2,7,11,15], 9)
    [0, 1]
    >>> s.twoSum([3,2,4], 6)
    [1, 2]
    >>> s.twoSum([3,3], 6)
    [0, 1]
    """

    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashMap: dict[int, int] = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashMap:
                return [hashMap[complement], i]
            # 注意这里容易遗漏, 并且 key 是数组元素, value 是数组索引
            hashMap[nums[i]] = i
        return []
