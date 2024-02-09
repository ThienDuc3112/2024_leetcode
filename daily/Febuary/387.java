class Solution {
    public int firstUniqChar(String s) {
        int[] appeared = new int[26];
        for (int i = 0; i < s.length(); i++) {
            appeared[s.charAt(i) - 'a']++;
        }
        for (int i = 0; i < s.length(); i++) {
            if (appeared[s.charAt(i) - 'a'] == 1)
                return i;
        }
        return -1;
    }
}