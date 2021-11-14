package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type priorityQueue []*ListNode

func (h priorityQueue) Len() int           { return len(h) }
func (h priorityQueue) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h priorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *priorityQueue) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *priorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	var pq priorityQueue
	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}

	newList := new(ListNode)
	tmp := newList
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		tmp.Next = node
		tmp = tmp.Next
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}

	return newList.Next
}

func MakeLists(array [][]int) []*ListNode {
	var lists []*ListNode
	for i := 0; i < len(array); i++ {
		list := ListNode{
			Val:  0,
			Next: nil,
		}
		tmp := &list
		for j := 0; j < len(array[i]); j++ {
			tmp.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			tmp.Next.Val = array[i][j]
			tmp = tmp.Next
		}
		lists = append(lists, list.Next)
	}
	return lists
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	fmt.Println()
}

func main() {
	array := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	lists := MakeLists(array)
	ret := mergeKLists(lists)
	PrintList(ret)
}
