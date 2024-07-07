package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	level := 0
	traversPre(root, &result, level)

	return result
}

func traversPre(node *TreeNode, s *[][]int, level int) {
	if node == nil {
		return
	}

	if len(*s) == level {
		*s = append(*s, []int{})
	}

	(*s)[level] = append((*s)[level], node.Val)
	traversPre(node.Left, s, level+1)
	traversPre(node.Right, s, level+1)
}

func levelOrderNoRecursion(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	//        1
	//       / \
	//      2   3
	//     / \ / \
	//    4  5 6  7
	// [ [1], [2,3], [4, 5, 6, 7] ]

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
