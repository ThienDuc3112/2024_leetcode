package removeNthFromEnd
class ListNode(var `val`: Int) {
    var next: ListNode? = null
}

class Solution {

    fun onePassRemoveNthFromEnd(head: ListNode?, n: Int): ListNode? {
        var leftPointer = head
        var rightPointer = head
        var m = n
        while (m > 0 && rightPointer != null) {
            rightPointer = rightPointer.next
            m--
        }
        if (rightPointer == null) return head?.next
        while (rightPointer != null && rightPointer.next != null) {
            leftPointer = leftPointer?.next
            rightPointer = rightPointer.next
        }
        leftPointer?.next = leftPointer?.next?.next
        return head
    }
    fun removeNthFromEnd(head: ListNode?, n: Int): ListNode? {
        var cur = head
        var length = 0
        while (cur != null) {
            cur = cur.next
            length++
        }
        length -= n
        if (length == 0) {
            return head?.next
        }
        cur = head
        while (length > 1 && cur != null && cur.next != null) {
            cur = cur.next
            length--
        }
        cur?.next = cur?.next?.next
        return head
    }
}
