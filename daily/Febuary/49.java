import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

// Group the anagram
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, List<String>> anagrams = new HashMap<>();
        for (String str : strs) {
            char[] temp = str.toCharArray();
            Arrays.sort(temp);
            String key = new String(temp);
            if (!anagrams.containsKey(key)) {
                anagrams.put(key, new ArrayList<>());
            }
            anagrams.get(key).add(str);
        }
        return new ArrayList<>(anagrams.values());
    }
}

// Inefficient af
// class Solution {
// public List<List<String>> groupAnagrams(String[] strs) {
// HashMap<Character, Integer>[] freqMaps = new HashMap[strs.length];
// for (int i = 0; i < strs.length; i++) {
// String str = strs[i];
// HashMap<Character, Integer> freqMap = new HashMap<>();
// for (int j = 0; j < str.length(); j++) {
// freqMap.put(str.charAt(j), freqMap.getOrDefault(str.charAt(j), 0) + 1);
// }
// freqMaps[i] = freqMap;
// }
// boolean[] checked = new boolean[strs.length];
// List<List<String>> res = new ArrayList<>();
// for (int i = 0; i < freqMaps.length; i++) {
// if(checked[i]) continue;
// List<String> anagrams = new ArrayList<>();
// anagrams.add(strs[i]);
// for(int j = i + 1; j < freqMaps.length; j++){
// if(freqMaps[i].equals(freqMaps[j])){
// anagrams.add(strs[j]);
// checked[j] = true;
// }
// }
// res.add(anagrams);
// }
// return res;

// }
// }