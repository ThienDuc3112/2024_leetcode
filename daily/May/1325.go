package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right != nil { root.Right = removeLeafNodes(root.Right, target)}
	if root.Left != nil { root.Left = removeLeafNodes(root.Left, target)}
    if root.Val == target && root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}
