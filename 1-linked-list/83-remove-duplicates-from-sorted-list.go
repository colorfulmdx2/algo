package linked_list

func deleteDuplicates(head *ListNode) *ListNode {
	fakeNode := &ListNode{0, head}
	iteratorNode := fakeNode
	hash := make(map[int]bool)

	for iteratorNode.Next != nil {
		// Check the value of the next node
		if _, ok := hash[iteratorNode.Next.Val]; ok {
			// If it's a duplicate, skip it
			iteratorNode.Next = iteratorNode.Next.Next
		} else {
			// If it's not a duplicate, add it to the hash map and move to the next node
			hash[iteratorNode.Next.Val] = true
			iteratorNode = iteratorNode.Next
		}
	}
	return fakeNode.Next
}
