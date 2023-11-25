from typing import List 

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashMap = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in hashMap:
                return [hashMap[complement], i]
            # 注意这里容易遗漏, 并且 key 是数组元素, value 是数组索引
            hashMap[nums[i]] = i
        return None