package main

func minimumDistance(nums []int) int {

	numSet := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := numSet[v]; !ok {
			numSet[v] = 1
			continue
		}
		numSet[v]++
	}
	smallest := 100000007
	for key, v := range numSet {
		if v < 3 {
			continue
		}
		i, j, k := -1,-1,-1
		for z := 0; z < len(nums); z++ {
			if nums[z] != key {
				continue
			}
			if i == -1 {
				i = z
				continue
			} else if j == -1 {
				j = z
				continue;
			} else if k == -1 {
				k = z
				smallest = min((j-i)+(k-j)+(k-i), smallest)
				continue
			}
			i, j, k = j, k, z
			smallest = min((j-i)+(k-j)+(k-i), smallest)
		} 
	}
	if smallest == 100000007 {
		return -1
	}
	return smallest
}