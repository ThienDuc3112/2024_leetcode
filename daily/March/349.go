package main

func intersection(nums1 []int, nums2 []int) []int {
    seen := map[int]bool{}
    for _, num := range(nums1) {
        seen[num] = false
    }
    res := make([]int, 0)
    for _, num := range(nums2) {
        if value, exist := seen[num]; exist && !value {
            res = append(res, num)
            seen[num] = true
        }
    }
    return res

    // Suboptimal code
    // set1 := map[int]bool{}
    // for _, num := range(nums1) {
    //     set1[num] = true
    // }
    // set2 := map[int]bool{}
    // for _, num := range(nums2) {
    //     set2[num] = true
    // }
    // if len(set1) > len(set2) {
    //     set1, set2 = set2, set1
    // }
    // res := []int{}
    // for k := range(set1) {
    //     if set2[k] == true {
    //         res = append(res, k)
    //     }
    // }
    // return res
}
