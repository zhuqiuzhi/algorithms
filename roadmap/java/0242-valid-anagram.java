package roadmap.java;

class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        int[] store = new int[26];
        for (int i = 0; i < s.length(); i++) {
            // 注意这里需要减去 a, 否则会报 java.lang.ArrayIndexOutOfBoundsException
            store[s.charAt(i)-'a']++;
            store[t.charAt(i)-'a']--;
        }

        for (int i =0; i < store.length; i++) {
            if (store[i] != 0) {
                return false; 
            }
        }

        return true;
    }
}