package linked_list

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]*ListNode)

	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		} else {
			hash[head] = head
			head = head.Next
		}
	}

	return false
}
