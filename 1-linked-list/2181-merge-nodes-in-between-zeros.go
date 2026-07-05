package linked_list

func mergeNodes(head *ListNode) *ListNode {
	fakeNode := &ListNode{}
	result := fakeNode

	for iterator, sum := head.Next, 0; iterator != nil; iterator = iterator.Next {
		if iterator.Val == 0 {
			fakeNode.Next = &ListNode{Val: sum, Next: nil}
			fakeNode = fakeNode.Next
			sum = 0
			continue
		}
		sum += iterator.Val
	}

	return result.Next
}
