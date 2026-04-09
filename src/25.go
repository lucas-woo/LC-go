package main

func removeDuplicates(nums []int) int {
	var ind int = 0;
	var before int;
	for i, v := range nums {
		if v != before || i ==0 {
			nums[ind] = v
			before = v;
			ind ++;
		}
	}
	return ind;
}