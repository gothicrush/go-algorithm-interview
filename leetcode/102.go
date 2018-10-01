/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type info struct {
    node *TreeNode
    level int
}

func newInfo(node *TreeNode, level int) info {
    return info {
        node: node,
        level: level,
    }
}

func levelOrder(root *TreeNode) [][]int {
 
    var ret [][]int
    
    if root == nil {
        return ret
    }
    
    var queue []info
    queue = append(queue, newInfo(root, 0))
    
    for len(queue) != 0 {
        
        node := queue[0].node
        level := queue[0].level
        
        queue = queue[1:]
        
        if level == len(ret) {
            ret = append(ret, []int{})
        }
        
        ret[level] = append(ret[level], node.Val)
        
        if node.Left != nil {
            queue = append(queue, newInfo(node.Left,level+1))
        }
        
        if node.Right != nil {
            queue = append(queue, newInfo(node.Right,level+1))
        }
    }
    
    return ret
}
