package main

func reorderList(head *ListNode)  {
    list := make([]*ListNode, 0)
    for head != nil {
        list = append(list, head)
        head = head.Next
    }
    left := 0
    right := len(list) - 1
    for left < right {
        list[left].Next = list[right]
        left++
        if left == right {
            break
        }
        list[right].Next = list[left]
        right--
    }
    list[left].Next = nil
}
