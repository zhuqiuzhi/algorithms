from typing import List

# Leetcode Problem: https://leetcode.com/problems/subsets/
# Given an integer array nums of unique elements, return all possible
# subsets (the power set).
# In mathematics, the power set of a set S is the set of all subsets of S,
# including the empty set and S itself.
#
# The solution set must not contain duplicate subsets. Return the solution in any order.
# 限制: 不能重复, 和排列有区别
# 技巧: 对 n 个数字依次做 0 和 1 选择
# 可以作为回溯算法的模板


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []

        subset = []

        def dfs(i: int):
            # base case
            if i == len(nums):
                res.append(subset.copy())
                return

            # descide to include nums[i]
            subset.append(nums[i])
            dfs(i+1)

            # descide not to include nums[i]
            subset.pop()
            dfs(i+1)

        dfs(0)

        return res
