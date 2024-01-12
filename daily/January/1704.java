class Solution {
    final String vowels = "aeiouAEIOU";

    boolean isVowel(char c) {
        for (int i = 0; i < vowels.length(); i++) {
            if (vowels.charAt(i) == c) {
                return true;
            }
        }
        return false;
    }

    public boolean halvesAreAlike(String s) {
        int left = 0;
        int right = 0;
        for (int i = 0; i < s.length(); i++) {
            if (isVowel(s.charAt(i))) {
                if (i < s.length() / 2)
                    left++;
                else
                    right++;
            }
        }
        return left == right;
    }
}