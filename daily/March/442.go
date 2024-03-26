package main

import "fmt"

// Need to check answer
func findDuplicates(nums []int) []int {
	for i := range nums {
		if nums[i] == i+1 || nums[i] < 0 {
			continue
		}
		val := nums[i]
		nums[i] = -2
		for val > 0 {
			fmt.Println(nums, val)
			if nums[val-1] == val {
				nums[val-1] = -1
				break
			}
			nums[val-1], val = val, nums[val-1]
		}
	}
	fmt.Println(nums)
	res := make([]int, 0)
	for i, v := range nums {
		if v == -1 {
			res = append(res, i+1)
		}
	}
	return res
}
