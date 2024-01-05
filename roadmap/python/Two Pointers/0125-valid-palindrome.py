class Solution:
    def isPalindrome(self, s: str) -> bool:
        """ Solution to https://leetcode.com/problems/valid-palindrome

        >>> Solution().isPalindrome("A man, a plan, a canal: Panama")
        True
        >>> Solution().isPalindrome("race a car")
        False
        >>> Solution().isPalindrome(" ")
        True
        """
        i, j = 0, len(s) - 1
        while i < j:
            if not s[i].isalnum():
                i = i + 1
                continue
            if not s[j].isalnum():
                j = j - 1
                continue

            if s[i].lower() != s[j].lower():
                return False

            i = i + 1
            j = j - 1
        return True

    def alphanum(self, c):
        return (
            ord("A") <= ord(c) <= ord("Z")
            or ord("a") <= ord(c) <= ord("z")
            or ord("0") <= ord(c) <= ord("9")
        )
