package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Should've just dfs instead
func sumOfLeftLeaves(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode, left bool)
    dfs = func(root *TreeNode, left bool) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            if left {
                res += root.Val
            }
            return
        }
        dfs(root.Left, true)
        dfs(root.Right, false)
    }
    dfs(root, false)
    return res
}
