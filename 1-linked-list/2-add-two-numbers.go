package linked_list

func AddTwoNumbers(firstList *ListNode, secondList *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}

	iterator := dummy

	carry := 0

	for firstList != nil || secondList != nil || carry != 0 {
		firstAddend := 0
		if firstList != nil {
			firstAddend = firstList.Val
			firstList = firstList.Next
		}

		secondAddend := 0
		if secondList != nil {
			secondAddend = secondList.Val
			secondList = secondList.Next
		}

		num := firstAddend + secondAddend + carry
		carry = num / 10
		iterator.Next = &ListNode{
			Val: num % 10,
		}
		iterator = iterator.Next
	}

	return dummy.Next
}
