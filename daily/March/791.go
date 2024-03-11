package main

import (
	"sort"
	"strings"
)

// Counting sort would've also work
func customSortString(order string, s string) string {
	var less func(i, j int) bool
	sb := strings.Split(s, "")
    less = func(i, j int) bool {
        priority1 := strings.Index(order, sb[i])
        priority2 := strings.Index(order, sb[j])
        if priority1 != -1 && priority2 != -1 {
            return priority1 < priority2
        } else if priority1 != -1 {
            return true
        } else if priority2 != -1 {
            return false
        }
        return sb[i] < sb[j]
	}
    sort.Slice(sb, less)
    return strings.Join(sb, "")
}
