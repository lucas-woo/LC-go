package main

func separateDigits(nums []int) []int {
	arr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		temp := make([]int, 0)
		for nums[i] > 0 {
			temp = append(temp, nums[i] % 10)
			nums[i] /= 10
		}
		for j := len(temp) - 1; j >= 0; j-- {
			arr = append(arr, temp[j])
		}
	}
	return arr
}