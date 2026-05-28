package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MinHeap) Push(item any) {
	*h = append(*h, item.(int))
}
func (h *MinHeap) Pop() any {
	t := *h
	last := t[len(t) - 1]
	*h = t[:len(t)-1]
	return last
}

func nthUglyNumber(n int) int {

	arr := make(MinHeap, 0)
	exist := make(map[int]bool, 0)
	heap.Push(&arr, 1)
	for i := 0; i < n - 1; i++ {
		last := heap.Pop(&arr).(int)
		if _, ok := exist[last*2]; !ok {
			heap.Push(&arr, last*2)
			exist[last*2] = true
		}
		if _, ok := exist[last*3]; !ok {
			heap.Push(&arr, last*3)
			exist[last*3] = true
		}
		if _, ok := exist[last*5]; !ok {
			heap.Push(&arr, last*5)
			exist[last*5] = true
		}
	}
	return heap.Pop(&arr).(int)
}