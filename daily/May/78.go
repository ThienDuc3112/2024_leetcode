package main

func subsets(nums []int) [][]int {
    res := make([][]int, 0)
    res = append(res, []int{})
    var help func(int)
    help = func(curIdx int) {
        if curIdx >= len(nums) {return}
        n := len(res)
        for i := 0; i < n; i++{
            set := res[i]
            copiedSet := make([]int, len(set) + 1)
            copy(copiedSet, set)
            copiedSet[len(set)] = nums[curIdx]
            res = append(res, copiedSet)
        }
        help(curIdx+1)
    }
    help(0)
    return res
}
