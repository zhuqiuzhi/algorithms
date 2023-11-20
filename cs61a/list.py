def min_abs_indies(s):
    """Indices of all elements in list s that the smallest absolute value

    >>> min_abs_indies([-4, -3, -2, 3, 2])
    [2, 4]
    >>> min_abs_indies([1, 2, 3, 4, 5])
    [0]
    """
    min_abs = min(map(abs, s))
    return [i for i in range(len(s)) if abs(s[i]) == min_abs]

def largest_adj_sum(s):
    """Largest sum of two adjacent elements in a list s.

    >>> largest_adj_sum([-4, -3, -2 , 3, 2, 4])
    6
    >>> largest_adj_sum([-4, 3, -2, -3, 2, -4])
    1
    """
    # range(-1) 会是空对象, [:-1] 表示到倒数第1个数为止
    return max([a+b for a, b in zip(s[:-1], s[1:])])
    # return max([s[i]+ s[i+1] for i in range(len(s)-1)])

def digit_dict(s):
    """Map each digit d to the lists of elements in s that end with d.

    >>> digit_dict([5, 8, 13, 21, 34, 55, 89])
    {1: [21], 3: [13], 4: [34], 5: [5, 55], 8: [8], 9: [89]}
    """ 
    last_digits = [ x%10 for x in s] 
    return {d: [ x for x in s if x%10 == d] for d in range(10) if d in last_digits }

def all_have_an_equal(s):
    """Does every element equal some other element in s?

    >>> all_have_an_equal([-4, -3, -2, 3, 2, 4])
    False
    >>> all_have_an_equal([4, 3, 2, 3, 2, 4])
    True
    """
    return all([s[i] in s[:i] + s[i+1:] for i in range(len(s))])

def fast_overlap(s, t, key=lambda x: x):
    """返回已排序无重复元素 list s 和 t 的相同元素的个数
    >>> s = [1, 4, 6, 9]
    >>> t = [4, 9]
    >>> fast_overlap(s, t)
    2
    """
    i, j, count = 0, 0, 0
    while i < len(s) and j < len(t):
        if s[i] == t[j]:
            count = count + 1
            i, j = i+1, j+1
        elif s[i] > t[j]:
            j = j+1
        else:
            i = i+1
    return count