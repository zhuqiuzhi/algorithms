class Solution:
    """
    >>> s = Solution()
    >>> s.isValid("()")
    True
    >>> s.isValid("()[]{}")
    True
    >>> s.isValid("[}")
    False
    >>> s.isValid("(")
    False
    >>> s.isValid(")")
    False
    """

    def isValid(self, s: str) -> bool:
        stack = []
        for i in range(len(s)):
            p = s[i]
            if p == "(" or p == "[" or p == "{":
                # push
                stack.append(p)
            elif p == ")" or p == "]" or p == "}":
                if len(stack) <= 0:
                    return False
                # pop
                left = stack.pop()
                if p == ")" and left != "(":
                    return False
                elif p == "]" and left != "[":
                    return False
                elif p == "}" and left != "{":
                    return False
        return len(stack) == 0


class Solution1(object):
    """
    >>> s = Solution1()
    >>> s.isValid("()")
    True
    >>> s.isValid("()[]{}")
    True
    >>> s.isValid("[}")
    False
    >>> s.isValid("(")
    False
    >>> s.isValid(")")
    False
    """

    def isValid(self, s):
        stack = []  # only use append and pop
        pairs = {
            '(': ')',
            '{': '}',
            '[': ']'
        }
        for bracket in s:
            if bracket in pairs:
                stack.append(bracket)
            elif len(stack) == 0 or bracket != pairs[stack.pop()]:
                return False

        return len(stack) == 0
