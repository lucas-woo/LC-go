package main

func leftRightDifference(nums []int) []int {
  left := make([]int, len(nums));
  right := make([]int, len(nums));
	cl := 0;
	cr := 0
	for i := 0; i < len(nums); i++ {
		left[i] = cl
		right[len(nums)-i-1] = cr
		cl += nums[i]
		cr += nums[len(nums)-i-1]
	}
	arr := make([]int, len(nums))
	for i := 0; i < len(arr); i++ {
		arr[i] = a(left[i] - right[i])
	}
	return arr
}
func a(n int) int {
	if n < 0 {
		return -n
	}
	return n
}