/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    var m map[int]int = make(map[int]int, 10)
    
    for pointer := head; pointer != nil; {
        m[pointer.Val]++
        pointer = pointer.Next
    }
    
    var dummyNode *ListNode = &ListNode{
        Next: head,
    }
    
    var slow *ListNode = dummyNode
    var fast *ListNode = head
    
    for fast != nil {
        
        if times, ok := m[fast.Val]; ok && times == 1 {
            slow.Next = fast
            slow = fast
        }
        
        fast = fast.Next
    }
    
    slow.Next = nil
    
    return dummyNode.Next
}
