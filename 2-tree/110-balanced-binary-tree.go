package tree

func isBalanced(root *TreeNode) bool {
	return dfs(root).balanced
}

type Result struct {
	balanced bool
	height   int
}

func dfs(node *TreeNode) Result {
	if node == nil {
		return Result{
			balanced: true,
			height:   0,
		}
	}

	left := dfs(node.Left)
	right := dfs(node.Right)
	balanced := left.balanced && right.balanced && abs(left.height-right.height) <= 1

	return Result{balanced, 1 + max(left.height, right.height)}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
