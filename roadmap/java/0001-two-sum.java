package roadmap.java;

import java.util.HashMap;
import java.util.Map;

class Solution {
     public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; i++) {
            int complement = target-nums[i];
            if (map.containsKey(complement)) {
                int[] result = { map.get(complement), i };
                return result;
                // return new int[]{map.get(complement), i};
            }
            map.put(nums[i], i);
        }

        return null;
    }
}
