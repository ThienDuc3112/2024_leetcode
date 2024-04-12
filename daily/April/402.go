package main

// Check answer
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	stack := make([]rune, 0)
	for _, c := range num {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > (c) {
			stack = stack[:len(stack)-1]
			k--
		}
		if len(stack) > 0 || c != '0' {
			stack = append(stack, c)
		}
	}
    for k > 0 && len(stack) > 0 {
        stack = stack[:len(stack) - 1]
        k--
    }
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}
