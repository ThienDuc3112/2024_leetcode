class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int nums1pointer = 0;
        int nums2pointer = 0;
        int[] res = new int[nums1.length];
        while (nums1pointer < m && nums2pointer < n) {
            if (nums1[nums1pointer] < nums2[nums2pointer]) {
                res[nums1pointer + nums2pointer] = nums1[nums1pointer];
                nums1pointer++;
            } else {
                res[nums1pointer + nums2pointer] = nums2[nums2pointer];
                nums2pointer++;
            }
        }
        while (nums1pointer < m) {
            res[nums1pointer + nums2pointer] = nums1[nums1pointer];
            nums1pointer++;
        }
        while (nums2pointer < n) {
            res[nums1pointer + nums2pointer] = nums2[nums2pointer];
            nums2pointer++;
        }
        for (int i = 0; i < res.length; i++) {
            nums1[i] = res[i];
        }
    }
}