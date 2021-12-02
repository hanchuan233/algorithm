package main

import "container/heap"

type PriorityQueue []int
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PriorityQueue) Pop() interface{} {
	val := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return val
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	sum := m + n
	if sum == 0 {
		return 0
	}

	pq := PriorityQueue{}
	cnt := sum / 2 + 1
	for _, val := range nums1 {
		heap.Push(&pq, val)
		if len(pq) > cnt {
			heap.Pop(&pq)
		}
	}
	for _, val := range nums2 {
		heap.Push(&pq, val)
		if len(pq) > cnt {
			heap.Pop(&pq)
		}
	}
	if sum % 2 == 0 {
		i := heap.Pop(&pq).(int)
		j := heap.Pop(&pq).(int)
		return float64(i + j) / 2
	} else {
		return float64(heap.Pop(&pq).(int))
	}
}
