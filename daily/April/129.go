package main

func sumNumbers(root *TreeNode) int {
    res := 0
    var dfs func (root *TreeNode, cur int)
    dfs = func(root *TreeNode, cur int) {
        if root == nil {
            return
        }
        val := cur * 10 + root.Val
        if root.Left == nil && root.Right == nil {
            res += val
            return
        }
        dfs(root.Left, val)
        dfs(root.Right, val)
    }
    dfs(root, 0)
    return res
}
