import java.util.HashMap;
import java.util.HashSet;

// Close but not correct attempt
class Solution {
    public String minWindow(String s, String t) {
        int left = 0, right = 0, resLeft = 0, resRight = s.length();
        boolean found = false;
        HashMap<Character, Integer> table = new HashMap<>();
        HashSet<Character> tSet = new HashSet<>();
        for (int i = 0; i < t.length(); i++) {
            table.put(t.charAt(i), table.getOrDefault(t.charAt(i), 0) + 1);
            tSet.add(t.charAt(i));
        }
        while (right < s.length()) {
            while (right < s.length() && !table.isEmpty()) {
                if (table.containsKey(s.charAt(right))) {
                    int newCount = table.get(s.charAt(right)) - 1;
                    if (newCount == 0)
                        table.remove(s.charAt(right));
                    else
                        table.put(s.charAt(right), newCount);
                }
                right++;
            }
            if (!table.isEmpty())
                break;
            else if (right - left < resRight - resLeft) {
                found = true;
                resLeft = left;
                resRight = right;
            }
            while (left < right && table.isEmpty()) {
                if (tSet.contains(s.charAt(left)))
                    table.put(s.charAt(left), 1);
                left++;
            }
        }
        if (found)
            return s.substring(resLeft, resRight);
        else
            return "";
    }
}

// Solution
// class Solution {
// public String minWindow(String s, String t) {
// if (s == null || t == null || s.length() ==0 || t.length() == 0 ||
// s.length() < t.length()) {
// return new String();
// }
// int[] map = new int[128];
// int count = t.length();
// int start = 0, end = 0, minLen = Integer.MAX_VALUE,startIndex =0;
// for (char c :t.toCharArray()) {
// map[c]++;
// }
// char[] chS = s.toCharArray();
// while (end < chS.length) {
// if (map[chS[end++]]-- >0) {
// count--;
// }
// while (count == 0) {
// if (end - start < minLen) {
// startIndex = start;
// minLen = end - start;
// }
// if (map[chS[start++]]++ == 0) {
// count++;
// }
// }
// }

// return minLen == Integer.MAX_VALUE? new String():
// new String(chS,startIndex,minLen);
// }
// }