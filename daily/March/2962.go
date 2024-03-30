package main

func countSubarrays(nums []int, k int) int64 {
	count := 0
	var res int64 = 0
	leftIdx := 0
	maxElement := 0
	for _, num := range nums {
		maxElement = max(maxElement, num)
	}
	for nums[leftIdx] != maxElement {
		leftIdx++
	}
	for _, right := range nums {
		if right == maxElement {
			count++
		}
		for count > k {
			if nums[leftIdx] == maxElement {
				count--
				leftIdx++
			}
			for nums[leftIdx] != maxElement {
				leftIdx++
			}
		}
		if count == k {
			res += int64(leftIdx + 1)
		}
	}
	return res
}
