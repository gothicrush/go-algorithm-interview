/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    
    var ret []int
    
    preorder(root, &ret)

    return ret
}

func preorder(node *TreeNode, retPointer *[]int) {
    
    if node == nil {
        return
    }
    
    *retPointer = append(*retPointer, node.Val)
    preorder(node.Left, retPointer)
    preorder(node.Right, retPointer)
}
