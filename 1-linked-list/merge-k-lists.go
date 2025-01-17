package linked_list

import (
	"container/heap"
	"fmt"
)

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}

	return dummy.Next
}

func main() {
	// Example usage:
	// Creating some test lists
	list1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	list3 := &ListNode{2, &ListNode{6, nil}}

	lists := []*ListNode{list1, list2, list3}

	// Merging k sorted lists
	result := mergeKLists(lists)

	// Printing the merged list
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}
