package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverspre(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}
	*s = append(*s, node.Val)
	traverspre(node.Left, s)
	traverspre(node.Right, s)
}

func preorderTraversal(root *TreeNode) []int {
	s := make([]int, 0)
	traverspre(root, &s)
	return s
}

func preorderTraversal2(root *TreeNode) []int {
	//        1
	//       / \
	//      2   3
	//     / \ / \
	//    4  5 6  7
	// 1 2 4 5 3 6 7
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	result := make([]int, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
