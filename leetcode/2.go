/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    for pt1,pt2 := l1,l2; pt1.Next != nil || pt2.Next != nil; {
        
        if pt1.Next == nil {
            pt1.Next = &ListNode{0, nil}
        }
        
        if pt2.Next == nil {
            pt2.Next = &ListNode{0, nil}
        }
        
        pt1 = pt1.Next
        pt2 = pt2.Next
    } 
    
    var res *ListNode = &ListNode{0, nil}
    var pt3 *ListNode = res
    var incr int = 0
    
    for pt1,pt2 := l1,l2; pt1 != nil && pt2 != nil; pt1,pt2,pt3 = pt1.Next,pt2.Next,pt3.Next{
         
        var sum int = pt1.Val + pt2.Val + incr
        if sum >= 10 {
            sum = sum - 10
            incr = 1
        } else {
            incr = 0
        }
        
        pt3.Next = &ListNode{sum, nil}
    }
    
    if incr != 0 {
        pt3.Next = &ListNode{1, nil}
    }
    
    return res.Next
}

