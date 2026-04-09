package main


func removeElement(nums []int, val int) int {
	var ind int = 0;
	for _, v := range nums {
		if v != val {
			nums[ind] = v
			ind++;
		}
	}
	return ind
}