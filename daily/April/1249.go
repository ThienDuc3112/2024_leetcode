package main

func minRemoveToMakeValid(s string) string {
    type StackItem struct {
        open bool
        index int
    }
    stack := make([]StackItem,0)
    for i, char := range s {
        if char == '(' {
            stack = append(stack, StackItem{true, i})
        } else if char == ')' {
            if len(stack) > 0 && stack[len(stack)-1].open {
                stack = stack[:len(stack)-1]
            } else {
                stack = append(stack, StackItem{false, i})
            }
        }
    }
    removed := 0
    for _, item := range stack {
        s = s[:item.index - removed] + s[item.index + 1 - removed:]
        removed++
    }
    return s
}
