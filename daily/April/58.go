package main

func lengthOfLastWord(s string) int {
    idx := len(s) - 1
    for idx > 0 {
        if s[idx] != ' ' {
            break
        }
        idx--
    }
    res := 0
    for idx >= 0 && s[idx] != ' ' {
        res++
        idx--
    }
    return res
}
