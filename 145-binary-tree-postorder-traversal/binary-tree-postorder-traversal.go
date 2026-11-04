/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res:= []int{}
    var inorder func (root *TreeNode)
    inorder = func (node * TreeNode){
        if node == nil{
            return
        }
        inorder(node.Left)
        inorder(node.Right)
        res = append(res,node.Val)
    }
    inorder(root)
    return res
}