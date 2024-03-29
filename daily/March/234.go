package main

func isPalindrome(head *ListNode) bool {
    arr := make([]int, 0)
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    left := 0
    right := len(arr) - 1
    for left < right {
        if arr[left] != arr[right] {return false}
        left++
        right--
    }
    return true
}
