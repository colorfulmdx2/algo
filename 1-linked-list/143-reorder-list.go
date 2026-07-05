package linked_list

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondPartHead := slow.Next
	slow.Next = nil // split list by 2

	return secondPartHead
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prevNext *ListNode

	for head != nil {
		prevNext, head, head.Next = head, head.Next, prevNext
	}

	return prevNext
}

func reorderList(head *ListNode) {
	firstPartHead := head
	lastPartHead := reverseLinkedList(middleNode(head))

	for lastPartHead != nil {
		firstListNextTmp, lastListNextTmp := firstPartHead.Next, lastPartHead.Next

		firstPartHead.Next = lastPartHead
		lastPartHead.Next = firstListNextTmp

		firstPartHead = firstListNextTmp
		lastPartHead = lastListNextTmp
	}
}
