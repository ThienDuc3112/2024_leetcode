package main

func removeNodes(head *ListNode) *ListNode {
    monoStack := make([]*ListNode, 0)
    node := head
    for node != nil {
        for len(monoStack) > 0 && monoStack[len(monoStack)-1].Val < node.Val {
            monoStack = monoStack[:len(monoStack)-1]
        }
        monoStack = append(monoStack, node)
        node = node.Next
    }
    var res *ListNode = nil
    cur := res
    for _, node := range monoStack {
        if cur == nil {
            res = node
            cur = node
        } else {
            cur.Next = node
            cur = cur.Next
        }
    }
    return res
}
