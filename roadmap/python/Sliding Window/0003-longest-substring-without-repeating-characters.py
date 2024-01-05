class Solution:
    """
    >>> s = Solution()
    >>> s.lengthOfLongestSubstring('abcabcbb')
    3
    >>> s.lengthOfLongestSubstring('bbbbb')
    1
    >>> s.lengthOfLongestSubstring('pwwkew')
    3
    """

    def lengthOfLongestSubstring(self, s: str) -> int:
        charSet = set()
        l = 0  # l 是滑动窗口的左边索引
        res = 0

        for r in range(len(s)):
            # 将滑动窗口 s[r] 左侧(包括自己)删除, 维持窗口内字符都是不
            # 注意这里是直到 s[r] 不在窗口元素集合 charSet 中才停止
            while s[r] in charSet:
                # 从集合中删除滑动窗口最左端的字符, 然后更新窗口左端索引
                # l 不停增大,直到 l 到达 s[r]字符所在位置
                charSet.remove(s[l])
                l += 1
            charSet.add(s[r])
            # update max size
            res = max(res, r-l+1)
        return res
