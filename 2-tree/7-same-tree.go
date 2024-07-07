package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	queue := []*TreeNode{p, q}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Left)
		queue = append(queue, left.Right, right.Right)
	}
	return true
}
