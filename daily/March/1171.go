package main

func removeZeroSumSublists(head *ListNode) *ListNode {
	res := head
	sum := 0
	nodeMap := make(map[int]*ListNode)
	for head != nil {
		sum += head.Val
		if sum == 0 {
			res = head.Next
			head = head.Next
            nodeMap = make(map[int]*ListNode)
			continue
		}
		node, ok := nodeMap[sum]
		if ok {
            removed := node.Next
            removeSum := removed.Val
			node.Next = head.Next
            for removeSum != 0 {
                delete(nodeMap, sum + removeSum)
                removed = removed.Next
                removeSum += removed.Val
            }
		} else {
			nodeMap[sum] = head
		}
		head = head.Next
	}
	return res
}
