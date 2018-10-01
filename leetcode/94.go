/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    
    var ret []int
    
    inorder(root, &ret)

    return ret
}

func inorder(node *TreeNode, retPointer *[]int) {
    
    if node == nil {
        return
    }
    
    inorder(node.Left, retPointer)
    *retPointer = append(*retPointer, node.Val)
    inorder(node.Right, retPointer)
}
