import java.util.Stack;
import java.util.Map;
import java.util.HashMap;

class Solution {
    public boolean isValid(String s) {
        Stack<Character> brackets = new Stack<>();
        Map<Character, Character> bracketLookup = new HashMap<>(3);

        bracketLookup.put('(', ')');
        bracketLookup.put('{', '}');
        bracketLookup.put('[', ']');

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);

            if (bracketLookup.containsKey(c)) {
                brackets.push(c);
            } else if (brackets.isEmpty() || bracketLookup.get(brackets.pop()) != c) {
                return false;
            }
        }

        return brackets.isEmpty();
    }
}