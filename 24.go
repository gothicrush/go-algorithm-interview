/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
 
    var dummyNode *ListNode = &ListNode {
        Val: 0,
        Next: head,
    }
    
    var prev *ListNode = dummyNode
    
    for prev.Next != nil && prev.Next.Next != nil {
        var n1 *ListNode = prev.Next
        var n2 *ListNode = n1.Next
        var next *ListNode = n2.Next
        
        
        n2.Next = n1
        n1.Next = next
        prev.Next = n2
        
        prev = n1
    }
    
    return dummyNode.Next
}
