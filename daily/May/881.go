package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	left := 0
	right := len(people) - 1
	res := 0
    sort.Ints(people)

	for right >= left {
		if people[right]+people[left] <= limit {
			left++
		}
		right--
		res++
	}

	return res
}
