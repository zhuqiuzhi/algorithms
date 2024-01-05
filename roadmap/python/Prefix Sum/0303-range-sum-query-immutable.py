from typing import List


class NumArray:
    """ Your NumArray object will be instantiated and called as such:
    obj = NumArray(nums)
    param_1 = obj.sumRange(left,right)

    >>> numArray = NumArray([-2, 0, 3, -5, 2, -1])
    >>> numArray.sumRange(0, 2)
    1
    >>> numArray.sumRange(2, 5)
    -1
    >>> numArray.sumRange(0, 5)
    -3
    """

    def __init__(self, nums: List[int]):
        self.prefixSums = []
        prefixSum = 0
        for num in nums:
            prefixSum = prefixSum + num
            self.prefixSums.append(prefixSum)

    def sumRange(self, left: int, right: int) -> int:
        # [0, right] - [0, left)
        assert (left >= 0 and left < len(self.prefixSums))
        assert (right < len(self.prefixSums))
        if left == 0:
            return self.prefixSums[right]
        return self.prefixSums[right] - self.prefixSums[left-1]
