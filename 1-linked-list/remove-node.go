package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeNode := &ListNode{0, head}      // фейковая нода, чтобы обработать кейс с пустым листом
	iteratorNode, length := fakeNode, 0 // переменная для итерации, длинна

	for iteratorNode.Next != nil { // итерируемся по листу, чтобы определить длинну
		length++
		iteratorNode = iteratorNode.Next
	}

	iteratorNode = fakeNode            // сбрасываем результат итерации до head (fake node)
	for i := 0; i <= length-n-1; i++ { // итерируемся до элемента с индексом length-n-1 (элеимент перед тем который нужно удалить)
		iteratorNode = iteratorNode.Next
	}

	iteratorNode.Next = iteratorNode.Next.Next // переписываем связь тем самым удаляя ноду

	return fakeNode.Next // возвращаем настоящую голову списка
}

func removeNthFromEnd2Pointers(head *ListNode, n int) *ListNode {
	fakeNode := &ListNode{0, head} // #1 - fake node (to handle case with empty list)

	slow, fast := fakeNode, fakeNode // #2 - two pointers for iteration

	for i := 0; i <= n; i++ { // #3 - move fast pointer to n (node we want to delete)
		fast = fast.Next
	}

	for fast != nil { //#4 - move both pointers in the loop while fast is not nil #4
		slow, fast = slow.Next, fast.Next
	}

	slow.Next = slow.Next.Next
	// #5
	// when fast is nil, slow pointer will be on the n-1 index of the list,
	// so we can rewrite connection of this node

	head, fakeNode = fakeNode.Next, nil // #6 remove fakeNode from the list, to not affect real data

	return head // #7 return real head of the list
}
