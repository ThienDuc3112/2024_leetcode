package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    slow := head
    if(slow == nil) {return false}
    fast := head.Next
    for fast != nil && slow != nil{
        if(slow == fast) {return true}
        fast = fast.Next
        if(fast == nil) {return false}
        fast = fast.Next
        slow = slow.Next

    }
    return false
}
