package tree

func traversino(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	traversino(node.Left, s)
	*s = append(*s, node.Val)
	traversino(node.Right, s)
}

func inorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)
	traversino(root, &s)
	return s
}

func inorderTraversal2(root *TreeNode) []int {
	//        1
	//       / \
	//      2   3
	//     / \ / \
	//    4  5 6  7
	// 4 2 5 1 6 3 7
	var result []int
	var stack []*TreeNode
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Val)

		current = current.Right
	}

	return result
}
