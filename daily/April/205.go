package main

func isIsomorphic(s string, t string) bool {
    mapping := make(map[byte]byte)
    set := make(map[byte]bool)
    for i := range s {
        target, mapped := mapping[s[i]]
        _, existed := set[s[i]]
        if (!mapped && existed) || (mapped && target != t[i]) {
            return false
        }
        mapping[s[i]] = t[i]
        set[t[i]] = true
    }
    return true
}
