package tree

import (
	"fmt"
)

// PostorderTraversal1 is the first implementation of post-order traversal.
func PostorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	//        1
	//       / \
	//      2   3
	//     / \ / \
	//    4  5 6  7

	//  4 5 2 6 7 3 1
	var result []int
	stack := []*TreeNode{root}
	resultReversed := []*TreeNode(nil)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		resultReversed = append(resultReversed, node)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	for i := len(resultReversed) - 1; i >= 0; i-- {
		result = append(result, resultReversed[i].Val)
	}

	return result
}

// PostorderTraversal2 is the second implementation of post-order traversal.
func PostorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var order []int

	s := []*TreeNode{root}

	for slen := len(s); slen > 0; slen = len(s) {
		top := s[slen-1]
		s = s[:slen-1]
		order = append([]int{top.Val}, order...)

		if top.Left != nil {
			s = append(s, top.Left)
		}

		if top.Right != nil {
			s = append(s, top.Right)
		}
	}

	return order
}

func logStack(stack []*TreeNode) {
	var arr []int
	for _, node := range stack {
		if node != nil {
			arr = append(arr, node.Val)
		}
	}
	fmt.Println(arr)
}
