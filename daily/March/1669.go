package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    headList2 := list2
    for list2.Next != nil {
        list2 = list2.Next
    }
    headList1 := list1
    tailList1 := list1
    i := 0
    for i < b {
        tailList1 = tailList1.Next
        i++
        if i == a-1 {
            headList1 = tailList1
        }
    }
    headList1.Next = headList2
    list2.Next = tailList1.Next
    return list1
}
