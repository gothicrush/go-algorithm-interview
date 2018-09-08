/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    
    var prev *ListNode = nil
    var cur *ListNode = head
    var next *ListNode = nil
    
    for cur != nil {
        next = cur.Next
        
        cur.Next = prev
        prev = cur
        cur = next
    }
    
    return prev
}
