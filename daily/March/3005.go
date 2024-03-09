package main
func maxFrequencyElements(nums []int) int {
    freqMap := map[int]int {}
    for i := 0; i < len(nums); i++ {
        val, ok := freqMap[nums[i]]
        if ok {
            freqMap[nums[i]] = val + 1
        } else {
            freqMap[nums[i]] = 1
        }
    }
    res := 0
    max := 0
    for _, v := range(freqMap) {
        if v == max {
            res += v
        } else if v > max {
            res = v
            max = v
        }
    }
    return res
}
