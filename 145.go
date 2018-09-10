/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    
    var ret []int
    
    postorder(root, &ret)

    return ret
    
}

func postorder(node *TreeNode, retPointer *[]int) {
    
    if node == nil {
        return
    }
    
    
    postorder(node.Left, retPointer)
    postorder(node.Right, retPointer)
    *retPointer = append(*retPointer, node.Val)

}
