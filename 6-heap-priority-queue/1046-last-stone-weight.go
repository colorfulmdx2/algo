package __heap_priority_queue

import "container/heap"

func lastStoneWeight(stones []int) int {
	pq := make(PriorityQueue, 0, len(stones))

	heap.Init(&pq)

	for pq.Len() > 1 {
		// достаём два самых тяжёлых
		x := heap.Pop(&pq).(int)
		y := heap.Pop(&pq).(int)
		if x != y {
			heap.Push(&pq, x-y) // кладём разницу
		}
	}

	if pq.Len() == 0 {
		return 0
	}

	return pq[0] // максимум (корень max-heap)
}
