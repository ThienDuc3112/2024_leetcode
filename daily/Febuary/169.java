import java.util.HashMap;

class Solution {

    public int betterMajorityElement(int[] nums) {
        int major = nums[0], count = 0;
        for (int i : nums) {
            if (count == 0) {
                major = i;
                count++;
            } else if (major == i)
                count++;
            else
                count--;
        }
        return major;
    }

    public int majorityElement(int[] nums) {
        HashMap<Integer, Integer> count = new HashMap<>();
        for (int i : nums)
            count.put(i, count.getOrDefault(i, 0) + 1);
        for (int i : count.keySet()) {
            if (count.get(i) > nums.length / 2)
                return i;
        }
        return -1;
    }
}