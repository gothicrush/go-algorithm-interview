/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    
    if head == nil {
        return head
    }
    
    var dummyNode *ListNode = &ListNode {
        Next: head,
    }
    
    var slow *ListNode = dummyNode
    var fast *ListNode = head
    
    for fast != nil {
        
        if fast.Val != val {
            slow.Next = fast
            slow = fast     
        }
        
        fast = fast.Next
    }
    
    slow.Next = nil
    
    return dummyNode.Next
}
