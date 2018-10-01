/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
 
    if root == nil {
        return 0
    }
    
    var leftHeight int = maxDepth(root.Left)
    var rightHeight int = maxDepth(root.Right)
    return getMaxHeight(leftHeight, rightHeight) + 1
}

func getMaxHeight(i, j int) int {
    
    if i < j {
        return j
    } else {
        return i
    }
}
