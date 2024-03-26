package main

// Technically solve by myself
// Helped by yesterday daily (which I checked the solution)
func firstMissingPositive(nums []int) int {
    for i := range nums {
        if i == nums[i] - 1 || nums[i] < 1 || nums[i] > len(nums) {
            continue
        }
        val := nums[i]
        for val > 0 && val <= len(nums) {
            val, nums[val-1] = nums[val-1], val
        }
    }
    res := 0
    for res = 0; res < len(nums); res++ {
        if res != nums[res] - 1 {
            break
        }
    }
    return res + 1
}
