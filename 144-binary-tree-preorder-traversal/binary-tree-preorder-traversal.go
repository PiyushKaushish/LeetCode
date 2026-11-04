/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res:= []int {}
    var pre func(root *TreeNode)
    pre = func(node *TreeNode){
        if node == nil{
            return 
        }
        res=append(res,node.Val)
        pre(node.Left)
        pre(node.Right)

    }
    pre(root)
    return res
}