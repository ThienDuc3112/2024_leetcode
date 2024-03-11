package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1] // || (intervals[i][1] == intervals[j][1] && intervals[i][0] > intervals[j][0])
    })
    res := 0
    time := intervals[0][1]
    for i := 0; i < len(intervals); i++ {
        if time > intervals[i][0] {
            res++
            continue
        }
        time = intervals[i][1]
    }
    return res
}
