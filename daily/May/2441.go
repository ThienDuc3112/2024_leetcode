package main

import "sort"

func findMaxK(nums []int) int {
	left := 0
	right := len(nums) - 1
	sort.Slice(nums, func(i, j int) bool {return nums[i] < nums[j]})
	for left < right {
		if -nums[left] == nums[right] {
			return nums[right]
		}
		if nums[left] > 0 || nums[right] < 0 {
			return -1
		}
        if -nums[left] > nums[right] {
            left++
        } else {
            right--
        }
	}
    return -1
}
