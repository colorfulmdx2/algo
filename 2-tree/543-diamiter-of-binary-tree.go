package tree

func diameterOfBinaryTree(root *TreeNode) int {
	depth := 0

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		depth = max(depth, left+right)

		return 1 + max(left, right)
	}

	dfs(root)

	return depth

}
