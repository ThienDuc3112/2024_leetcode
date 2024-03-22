package main

import "math"

// First solution was linear search

// O(log(n)) sol
func pivotInteger(n int) int {
	sum := n * (n + 1) / 2
	low, high := 1, n + 1
    for low < high {
        mid := (low + high) / 2
        leftSum := mid * (mid + 1) / 2
        rightSum := sum + mid - leftSum
        if leftSum == rightSum {return mid}
        if leftSum < rightSum {
            low = mid + 1
        } else {
            high = mid
        }
    }
	return -1
}

// Constant time solution
func constantPivotInteger(n int) int {
    x := math.Sqrt(float64(n*n + n)/2)
    if x == math.Floor(x) {
        return int(x)
    }
    return -1
}
