import java.util.HashSet;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        HashSet<Character> appeared = new HashSet<>();
        int max = 0;
        int left = 0;
        for (int right = 0; right < s.length(); right++) {
            char c = s.charAt(right);
            while (appeared.contains(c)) {
                appeared.remove(s.charAt(left));
                left++;
            }
            appeared.add(c);
            max = Math.max(right - left + 1, max);
        }
        return max;
    }
}