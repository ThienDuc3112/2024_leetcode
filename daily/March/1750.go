package main

func minimumLength(s string) int {
    left := 0
    right := len(s) - 1
    for left < right && s[left] == s[right] {
        curChar := s[left]
        for left <= right && s[left] == curChar {
            left++
        }
        for left <= right && s[right] == curChar {
            right--
        }
    }
    return max(0, right - left + 1)
}
