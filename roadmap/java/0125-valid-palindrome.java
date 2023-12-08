class Solution {
    public boolean isPalindrome(String s) {
        int i = 0;
        int j = s.length() - 1;
        // 不能是 i != j, 可能 j-1 不等于 i+1, 例如 s 是 "aa"
        while (i < j) {
            char start = s.charAt(i);
            // 忽略非字母或者数字的字符
            if (!Character.isLetterOrDigit(start)) {
                i++;
                continue;
            }
            char end = s.charAt(j);
            if (!Character.isLetterOrDigit(end)) {
                j--;
                continue;
            }

            // 转换成小写字母
            if (Character.toLowerCase(start)!= Character.toLowerCase(end)) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
}
