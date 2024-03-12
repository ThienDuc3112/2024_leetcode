package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool {return potions[i] < potions[j]})
    
}
