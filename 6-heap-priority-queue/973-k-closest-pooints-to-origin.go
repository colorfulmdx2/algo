package __heap_priority_queue

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	pq := &priorityQueuePoints{}

	heap.Init(pq)

	for _, point := range points {
		dist := point[0]*point[0] + point[1]*point[1]

		heap.Push(pq, Point{coords: point, distance: dist})

		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	res := make([][]int, 0, k)
	for _, v := range *pq {
		res = append(res, v.coords)
	}

	return res

}

type Point struct {
	coords   []int
	distance int
}

type priorityQueuePoints []Point

func (pq priorityQueuePoints) Len() int {
	return len(pq)
}

func (pq priorityQueuePoints) Less(i, j int) bool {
	return pq[i].distance > pq[j].distance // < min heap, with > it's maxheap
}

func (pq priorityQueuePoints) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueuePoints) Push(x interface{}) {
	*pq = append(*pq, x.(Point))
}

func (pq *priorityQueuePoints) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (h priorityQueuePoints) Top() Point { return h[0] }
