package __heap_priority_queue

import "container/heap"

type KthLargest struct {
	k  int
	pq PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
	}

	for _, x := range nums {
		kl.Add(x)
	}

	return kl
}

func (kl *KthLargest) Add(x int) int {
	if len(kl.pq) < kl.k {
		heap.Push(&kl.pq, x)
		return kl.pq.Top()
	}
	if x > kl.pq.Top() {
		heap.Pop(&kl.pq)
		heap.Push(&kl.pq, x)
	}
	return kl.pq.Top()
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j] // min heap, with > it's maxheap
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (h PriorityQueue) Top() int { return h[0] }

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
