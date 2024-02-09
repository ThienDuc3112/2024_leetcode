import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class Solution {
    // After checking answer
    public String frequencySort(String s) {
        HashMap<Character, Integer> freqMap = new HashMap<>();
        for (char c : s.toCharArray())
            freqMap.put(c, freqMap.getOrDefault(c, 0) + 1);
        List<Character>[] bucket = new List[s.length() + 1];
        for (char key : freqMap.keySet()) {
            int value = freqMap.get(key);
            if (bucket[value] == null)
                bucket[value] = new ArrayList<>();
            bucket[value].add(key);
        }
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < bucket.length; i++) {
            List<Character> list = bucket[i];
            if (list != null) {
                for (char character : list) {
                    for (int j = 0; j < i; j++) {
                        res.append(character);
                    }
                }
            }
        }
        return res.toString();
    }

    // Inefficient method
    // public String frequencySort(String s) {
    // HashMap<Character, Integer> freqMap = new HashMap<>();
    // for (int i = 0; i < s.length(); i++) {
    // freqMap.put(s.charAt(i), freqMap.getOrDefault(s.charAt(i), 0) + 1);
    // }
    // List<Map.Entry<Character, Integer>> listOfChars = new ArrayList<>();
    // for (Map.Entry<Character, Integer> entry : freqMap.entrySet()) {
    // listOfChars.add(entry);
    // }
    // listOfChars.sort((a, b) -> {
    // return a.getValue() - b.getValue();
    // });
    // String res = new String();
    // for (Map.Entry<Character, Integer> entry : listOfChars) {
    // for (int i = 0; i < entry.getValue(); i++) {
    // res += entry.getKey();
    // }
    // }
    // return res;
    // }
}