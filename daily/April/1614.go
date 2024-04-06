package main

func maxDepth(s string) int {
    res := 0
    cur := 0
    for _, char := range s {
        if char == '(' {
            cur++
        } else if char == ')' {
            cur--
        }
        res = max(cur, res)
    }
    return res
}
