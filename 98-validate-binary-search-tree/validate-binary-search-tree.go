/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    var validate func( root *TreeNode,min,max int) bool
    validate = func(node *TreeNode, min,max int) bool{
        if node == nil{
            return true
        }

        if node.Val <= min || node.Val >= max {
            return false
        }

        return (validate(node.Left,min,node.Val) && validate(node.Right,node.Val,max))
    }
    return validate(root,math.MinInt,math.MaxInt)
}