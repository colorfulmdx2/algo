package __array

import (
	"container/heap"
)

type Row struct {
	soldiers int
	index    int
}

// Define a PriorityQueue for the min-hea
type PriorityQueue []*Row

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].soldiers == pq[j].soldiers {
		return pq[i].index < pq[j].index
	}
	return pq[i].soldiers < pq[j].soldiers
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Row))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func kWeakestRows(mat [][]int, k int) []int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i, row := range mat {
		soldiers := 0
		for _, val := range row {
			if val == 1 {
				soldiers++
			} else {
				break
			}
		}
		heap.Push(pq, &Row{soldiers: soldiers, index: i})
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		row := heap.Pop(pq).(*Row)
		result[i] = row.index
	}

	return result
}
