import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    public List<List<String>> suggestedProducts(String[] products, String searchWord) {
        Arrays.sort(products);
        List<List<String>> res = new ArrayList<>();
        res.add(Arrays.asList(products));
        for (int i = 0; i < searchWord.length(); i++) {
            List<String> cur = new ArrayList<>();
            for (String str : res.get(res.size() - 1))
                if (str.length() > i && str.charAt(i) == searchWord.charAt(i))
                    cur.add(str);
            res.add(cur);
        }
        res.remove(0);
        res.replaceAll((list) -> {
            if (list.size() <= 3)
                return list;
            return list.subList(0, 3);
        });
        return res;
    }
}