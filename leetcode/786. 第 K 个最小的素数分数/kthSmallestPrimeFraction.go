package main

import (
	"container/heap"
	"fmt"
)

type fraction [2]int
type PriorityQueue []fraction
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i][0]*pq[j][1] > pq[j][0]*pq[i][1] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(fraction))
}
func (pq *PriorityQueue) Pop() interface{} {
	f := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return f
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	var pq PriorityQueue
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(&pq, fraction{arr[i], arr[j]})
			if pq.Len() > k {
				heap.Pop(&pq)
			}
		}
	}

	min := heap.Pop(&pq).(fraction)

	return []int{min[0], min[1]}
}

func main() {
	arr := []int{1, 2, 3, 5}
	ret := kthSmallestPrimeFraction(arr, 3)
	fmt.Println(ret)
}
