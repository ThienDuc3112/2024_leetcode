package main

import "strings"

func reversePrefix(word string, ch byte) string {
	var sb strings.Builder
    reversed := false
    prefixArr := make([]string, 0)
    prefix := ""

    for _, c := range word {
        sb.WriteRune(c)
        if c == rune(ch) && !reversed {
            prefixArr = strings.Split(sb.String(), "")
            for i, j := 0, len(prefixArr)-1; i < j; i, j = i+1, j-1 {
                prefixArr[i], prefixArr[j] = prefixArr[j], prefixArr[i]
            }
            prefix = strings.Join(prefixArr, "")
            sb.Reset()
            reversed = true
        }
    }
    return prefix + sb.String()
}
