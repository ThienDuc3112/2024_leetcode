class Solution {
    public String firstPalindrome(String[] words) {
        for (String string : words) {
            int left = 0, right = string.length() - 1;
            while (left < right) {
                if (string.charAt(left) != string.charAt(right))
                    break;
                left++;
                right--;
            }
            if (left >= right)
                return string;
        }
        return "";
    }
}