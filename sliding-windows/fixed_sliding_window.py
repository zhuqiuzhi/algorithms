from typing import List 

def fixed_sliding_window(arr: List[int], k: int) -> List[int]:
    """Get list of sum up k element subarray in arr
    >>> fixed_sliding_window([1, 2, 3, 4, 5, 6], 3) 
    [6, 9, 12, 15]
    >>> fixed_sliding_window([-3, 4, 2, 7, 5], 3)
    [3, 13, 14]
    """
    sum_subarray = sum(arr[:k])
    result = [sum_subarray]

    for i in range(1, len(arr)-k+1):
        mid_sum =  sum_subarray - arr[i-1]
        sum_subarray = mid_sum + arr[i+k-1]
        result.append(sum_subarray)
    
    return result
    