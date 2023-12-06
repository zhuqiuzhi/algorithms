# https://leetcode.com/problems/corporate-flight-bookings/
from typing import List

class Solution:
    """There are n flights that are labeled from 1 to n.
    You are given an array of flight bookings bookings, 
    where bookings[i] = [firsti, lasti, seatsi] represents a booking for flights firsti 
    through lasti (inclusive) with seatsi seats reserved for each flight in the range.
    Return an array answer of length n, where answer[i] 
    is the total number of seats reserved for flight i.

    >>> s = Solution() 
    >>> s.corpFlightBookings([[1,2,10],[2,3,20],[2,5,25]], 5)
    [10, 55, 45, 25, 25]
    >>> s.corpFlightBookings([[1,2,10],[2,2,15]], 2)
    [10, 25]
    """
    def corpFlightBookings(self, bookings: List[List[int]], n: int) -> List[int]:
        diffs = [0 for i in range(n)]
        answers = [0 for i in range(n)]

        # init diff array
        for booking in bookings:
            first = booking[0]-1
            last = booking[1]-1
            seats = booking[2]
            diffs[first] = diffs[first] + seats 
            if last + 1 < n:
                diffs[last+1] = diffs[last+1] - seats
        
        for i in range(0, n):
            if i == 0:
                answers[i] = diffs[i]
            else:
                # key point
                answers[i] = answers[i-1] + diffs[i]

        return answers
        