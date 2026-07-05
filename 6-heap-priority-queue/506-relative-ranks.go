package __heap_priority_queue

import (
	"container/heap"
	"slices"
	"strconv"
)

func FindRelativeRanks(score []int) []string {
	res := make([]string, len(score), len(score))
	pq := priorityQueue(append([]int(nil), score...))

	heap.Init(&pq)

	for i := 0; i < len(score); i++ {
		n := heap.Pop(&pq).(int)
		index := slices.Index(score, n)

		if i == 0 {
			res[index] = "Gold Medal"
			continue
		}

		if i == 1 {
			res[index] = "Silver Medal"
			continue
		}

		if i == 2 {
			res[index] = "Bronze Medal"
			continue
		}

		res[index] = strconv.Itoa(i + 1)
	}
	return res
}

type priorityQueue []int

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j] // < min heap, with > it's maxheap
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (h priorityQueue) Top() int { return h[0] }
