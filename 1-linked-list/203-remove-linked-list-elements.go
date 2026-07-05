package linked_list

func removeElements(head *ListNode, val int) *ListNode {
	fakeNode := &ListNode{0, head}
	iteratorNode := fakeNode

	for iteratorNode.Next != nil {
		if iteratorNode.Next.Val == val {
			iteratorNode.Next = iteratorNode.Next.Next
		} else {
			iteratorNode = iteratorNode.Next
		}
	}
	return fakeNode.Next
}
