/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    if head == nil {
        return head
    }
    
    var slow *ListNode = head
    var fast *ListNode = head
    
    for fast != nil {
        
        if slow.Val != fast.Val {
            slow.Next = fast
            slow = fast
        }
        
        fast = fast.Next
    }
    
    slow.Next = nil
    
    return head
}
