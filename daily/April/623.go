package main

func rAddOneRow(root *TreeNode, val int, depth int, left bool) *TreeNode {
    if depth <= 1 {
        newRoot := &TreeNode{Val: val, Left: nil, Right: nil}
        if left {
            newRoot.Left = root
        } else {
            newRoot.Right = root
        }
        return newRoot
    }
    if root == nil {
        return nil
    }
    root.Left = rAddOneRow(root.Left, val, depth - 1, true)
    root.Right = rAddOneRow(root.Right, val, depth - 1, false)
    return root
}
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    return rAddOneRow(root, val, depth, true)
}
