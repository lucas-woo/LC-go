package main

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len () int {
	return len(h)
}

func (h MaxHeap) Less (i int, j int) bool {
	return h[i] > h[j];
}

func (h MaxHeap) Swap (i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push (item any) {
	*h = append(*h, item.(int))
}
func (h *MaxHeap) Pop () any {
	x := (*h)[len((*h))-1]
	*h = (*h)[:len((*h))-1]
	return x
}

func lastStoneWeight(stones []int) int {
	ourHeap := &MaxHeap{}
	heap.Init(ourHeap)
	for _, v := range stones {
		heap.Push(ourHeap, v);
	}
	for ourHeap.Len() > 1 {
		big := heap.Pop(ourHeap).(int);
		small := heap.Pop(ourHeap).(int);
		if big > small {
			heap.Push(ourHeap, (big-small))
		}
	}
	if ourHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(ourHeap).(int)
}