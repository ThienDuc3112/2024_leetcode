package main

// This was from the "If programming was an anime" :skull:
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for true {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}
