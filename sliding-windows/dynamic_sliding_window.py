from typing import List
import sys

def dynamic_siding_window(arr: List[int], target: int) -> int:
    """

    >>> dynamic_siding_window([1, 2, 3, 4, 5, 6], 7)
    2
    >>> dynamic_siding_window([1, 3, 7, 9, 10], 7)
    1
    >>> dynamic_siding_window([1, 2, 4, 5, 6], 7)
    2
    """
    min_length = sys.maxsize 
    start , end = 0, 0
    current_sum= 0

    while end < len(arr):
       current_sum = current_sum + arr[end]

       while start <= end and current_sum >= target:
           min_length = min(min_length, end-start+1)

           if start + 1 <= end:
               current_sum = current_sum - arr[start]
               start = start + 1
           else:
               break

       end = end + 1

    return min_length