package linked_list

func removeNodes(head *ListNode) *ListNode {
	maxNode := head
	newHead := head
	curr := head
	iterator := head

	for iterator.Next != nil {
		if maxNode.Val > iterator.Val {
			curr = curr.Next
			newHead = curr.Next
			curr.Next = nil
			iterator = iterator.Next
			continue
		}

		if iterator.Val > curr.Val {
			maxNode = iterator
			newHead = curr.Next
			curr = curr.Next

			curr.Next = nil
			iterator = iterator.Next
			continue
		}
		curr = curr.Next
		iterator = iterator.Next
	}

	return newHead
}

func reverse(head *ListNode) *ListNode {
	var newNext *ListNode
	for head != nil {
		currentNext := head.Next
		head.Next = newNext
		newNext = head
		head = currentNext

	}
	return newNext
}
