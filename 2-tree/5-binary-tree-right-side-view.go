package tree

func binaryTreeRightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	//        1  <-
	//       / \
	//      2   3  <-
	//     / \ / \
	//    4  5 6  7  <-
	//      1 3 7

	var result []int
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
		level = level[levelSize-1:]
		result = append(result, level[0])
	}

	return result
}
