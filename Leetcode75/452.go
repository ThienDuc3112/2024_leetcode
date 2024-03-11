package main

import "sort"

// Checked answer
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i,j int) bool {
        return points[i][1] < points[j][1]
    })
    res := 1
    arrow := points[0][1]
    for i := 0; i < len(points); i++ {
        if arrow >= points[i][0] {continue}
        res++
        arrow = points[i][1]
    }
    return res
}
