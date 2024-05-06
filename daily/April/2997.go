package main

func minOperations(nums []int, k int) int {
    xor := 0
    for _, num := range nums {
        xor ^= num
    }
    
    res := 0
    for i := 0; i < 32; i++ {
        if (1 << i) & xor != (1 << i) & k {
            res++
        }
    }
    return res
}
