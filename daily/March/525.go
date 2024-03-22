package main

func findMaxLength(nums []int) int {
    first := map[int]int {
        0: -1,
    }
    count := 0
    res := 0
    for i, val := range nums {
        if val == 0 {
            count--
        } else {
            count++
        }
        if val, ok := first[count]; ok {
            res = max(res, i - val)
        } else {
            first[count] = i
        }
    }
    return res
}
