from typing import List

class Solution:
    """
    >>> s = Solution()
    >>> s.carPooling([[2,1,5],[3,3,7]], 4)
    False
    >>> s.carPooling([[2,1,5],[3,3,7]], 5)
    True
    """
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        diffs = [0 for i in range(0, 1001)] 
        stations = [0 for i in range(0, 1001)] 

        for trip in trips:
            numPassengersi = trip[0]
            begin = trip[1]
            end = trip[2] - 1
            diffs[begin] = diffs[begin] + numPassengersi
            if end + 1 < capacity:
                diffs[end+1] = diffs[end+1] - numPassengersi

        stations[0] = diffs[0]
        for i in range(1, 1001):
            stations[i] = stations[i-1] + diffs[i]

        return not any([x > capacity for x in stations])