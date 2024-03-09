class Solution {
    public String maximumOddBinaryNumber(String s) {
        int ones = 0;
        int zeros = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '1')
                ones++;
            else
                zeros++;
        }
        StringBuilder res = new StringBuilder("");
        while (ones > 1) {
            res.append('1');
            ones--;
        }
        while (zeros > 0) {
            res.append('0');
            zeros--;
        }
        res.append('1');
        return res.toString();
    }
}
