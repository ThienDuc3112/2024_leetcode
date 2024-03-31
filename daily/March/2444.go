package main

// Answer was checked :(
func countSubarraysHard(nums []int, minK int, maxK int) int64 {
	maxIdx := -1
	minIdx := -1
    badIdx := -1
	var res int64 = 0
	for right := range nums {
		if nums[right] > maxK || nums[right] < minK {
            badIdx = right
		} else if nums[right] == maxK {
            maxIdx = right
        } else if nums[right] == minK {
            minIdx = right
        }
        res += int64(max(min(maxIdx, minIdx) - badIdx, 0))
	}
	return res
}
