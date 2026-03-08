package main

import "math/big"
import "math"
func smallestBalancedIndex(nums []int) int {
	left := make([]int64, len(nums))
	right := make([]int64, len(nums))

	for i := 1; i < len(nums); i++ {
		left[i] = left[i - 1] + int64(nums[i - 1])
	}
	right[len(nums) - 1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
    a := big.NewInt(right[i + 1])
    b := big.NewInt(int64(nums[i + 1]))
    result := new(big.Int)
		result.Mul(a, b)		
		biggest := big.NewInt(math.MaxInt64)
		if x := result.Cmp(biggest); x == 1 {
			right[i] = -1
		} else {
			right[i] = right[i + 1] * int64(nums[i + 1])
		}
	}
	for i := 0; i < len(nums); i++ {
		if left[i] == right[i] {
			return i
		}
	}
	return -1
}
