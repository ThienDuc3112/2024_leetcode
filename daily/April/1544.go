package main

import (
	"strings"
)

func isSame(a, b byte) bool {
	return a-'a' == b-'A' || a-'A' == b-'a'
}

func makeGood(s string) string {
	var res strings.Builder
	good := false
	for !good {
        good = true
		for i := 0; i < len(s); i++ {
			if i <= len(s) - 2 && isSame(s[i], s[i+1]) {
				i++
				good = false
			} else {
				res.WriteByte(s[i])
			}
		}
        s = res.String()
        res.Reset()
	}
    return s
}
