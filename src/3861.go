package main

func minimumIndex(capacity []int, itemSize int) int {
	smallest := -1
	smallestIndex := -1
  for i := 0; i < len(capacity); i++ {
		if capacity[i] >= itemSize {
			smallest = capacity[i];
			smallestIndex = i
			break
		}
	}
  for i := 0; i < len(capacity); i++ {
		if capacity[i] >= itemSize && capacity[i] < smallest {
			smallest = capacity[i];
			smallestIndex = i
		}
	}
	return smallestIndex
}