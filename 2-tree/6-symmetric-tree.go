package tree

func symmetricTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	//        1
	//       / \
	//      2   2
	//     / \ / \
	//    3  4 4  3
	//

	queue := []*TreeNode{root.Left, root.Right}

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

		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}
	return true
}
