package main

// Needed solution =(
func checkValidString(s string) bool {
    maxNeeded := 0
    minNeeded := 0
    for _, c := range s {
        if c == '(' {
            maxNeeded++
            minNeeded++
        } else if c == ')' {
            maxNeeded--
            minNeeded = max(minNeeded - 1, 0)
        } else {
            maxNeeded++
            minNeeded = max(minNeeded - 1, 0)
        }
        if maxNeeded < 0 {
            return false
        }
    }
    return minNeeded == 0
}
