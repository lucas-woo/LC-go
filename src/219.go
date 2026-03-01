package main


func containsNearbyDuplicate(nums []int, k int) bool {
	var numMap map[int]int = make(map[int]int, 0)

	for i, v := range nums {
		if prevIndex, ok := numMap[v]; ok && i - prevIndex <= k {
			return true
		}
		numMap[v] = i
	}
	return false
}
