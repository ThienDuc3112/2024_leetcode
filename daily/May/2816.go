package main

func doubleAndCarryOver(node *ListNode) bool {
    if node == nil {return false}
    carry := doubleAndCarryOver(node.Next)
    res := false
    node.Val *= 2
    if carry {
        node.Val++
    }
    if node.Val >= 10 {
        res = true
        node.Val -=10
    }
    return res
}

func doubleIt(head *ListNode) *ListNode {
    carry := doubleAndCarryOver(head)
    if carry {
        head = &ListNode{Val: 1, Next: head}
    }
    return head
}
