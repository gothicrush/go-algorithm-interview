/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
 
    
    var plist []int = preorder(p)
    var qlist []int = preorder(q)
    
    if len(plist) != len(qlist) {
        return false
    }
    
    for i := 0; i < len(plist) ; i++ {
        if plist[i] != qlist[i] {
            return false
        }
    }
    
    return true
}

func preorder(p *TreeNode) []int {
    
    var ret []int
    
    preorderRecursion(p, &ret)
    
    return ret
    
}

func preorderRecursion(p *TreeNode, pret *[]int) {
    
    if p == nil {
        *pret = append(*pret, -1)
        return
    }
    
    *pret = append(*pret, p.Val)
    preorderRecursion(p.Left, pret)
    preorderRecursion(p.Right, pret)
}
