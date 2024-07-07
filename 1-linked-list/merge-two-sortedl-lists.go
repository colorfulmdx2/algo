package linked_list

import "math"

func Value(node *ListNode) float64 {
	if node == nil {
		return math.Inf(1)
	}
	return float64(node.Val)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	fakeNode := &ListNode{Val: 0}
	currNode := fakeNode

	for list1 != nil && list2 != nil {
		if Value(list1) < Value(list2) {
			currNode.Next = list1
			list1 = list1.Next
		} else {
			currNode.Next = list2
			list2 = list2.Next
		}
		currNode = currNode.Next
	}

	if list1 != nil {
		currNode.Next = list1
	} else {
		currNode.Next = list2
	}

	return fakeNode.Next
}
