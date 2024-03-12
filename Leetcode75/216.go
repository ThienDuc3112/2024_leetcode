package main

func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0)
    var backtrack func(int, int, int, []int) 
    backtrack = func(startFrom, maximum, remainingCount int, sequence []int) {
        if remainingCount == 0 {
            if maximum == 0 {
                duplicate := make([]int, len(sequence))
                copy(duplicate, sequence)
                res = append(res, duplicate)
            }
            return
        }
        for i := startFrom + 1; i <= maximum && i <= 9; i++ {
            sequence = append(sequence, i)
            backtrack(i, maximum - i, remainingCount - 1, sequence)
            sequence = sequence[:len(sequence)-1]
        }
    }
    backtrack(0, n, k, make([]int, 0))
    return res
}
