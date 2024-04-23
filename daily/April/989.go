package main

func reverse[T any](slice []T) {
    for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
        slice[i], slice[j] = slice[j], slice[i]
    }
}

func smallestFromLeaf(root *TreeNode) string {
	inserted := false
	res := ""
	curByteArr := make([]byte, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		curByteArr = append(curByteArr, 'a'+byte(root.Val))
		if root.Left == nil && root.Right == nil {
            reverse(curByteArr)
			if !inserted {
				inserted = true
				res = string(curByteArr)
			} else {
				res = min(res, string(curByteArr))
			}
            reverse(curByteArr)
			curByteArr = curByteArr[:len(curByteArr)-1]
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		curByteArr = curByteArr[:len(curByteArr)-1]
	}
	dfs(root)
	return res
}
