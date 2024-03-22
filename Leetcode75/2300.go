package main

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool { return potions[i] < potions[j] })
	res := make([]int, len(spells))
	var binarySearch func(int, int, int) int
	binarySearch = func(spell, left, right int) int {
		if left == right {
			return left
		}
		mid := (left + right) / 2
		var effectiveness int64 = int64(spell) * int64(potions[mid])
		if effectiveness >= success {
			return binarySearch(spell, left, mid)
		} else {
			return binarySearch(spell, mid+1, right)
		}
	}
	for i, spell := range spells {
		res[i] = len(potions) - binarySearch(spell, 0, len(potions))
	}
	return res
}
