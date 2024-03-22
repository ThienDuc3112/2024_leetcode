package main

// Did once before
func insert(intervals [][]int, newInterval []int) [][]int {
    res := make([][]int, 0)
    for _, interval := range intervals {
        if interval[1] < newInterval[0] {
            res = append(res, interval)
        } else if newInterval[1] < interval[0] {
            res = append(res, newInterval)
            newInterval = interval
        } else {
            newInterval = []int {min(newInterval[0], interval[0]), max(newInterval[1], interval[1])}
        }
    }
    res = append(res, newInterval)
    return res;
}
