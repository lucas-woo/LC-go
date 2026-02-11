package main

func longestBalanced(nums []int) int {
	longest := 0
	for i := 0; i < len(nums); i++ {
		seen := make(map[int]bool, 0)
		even := 0;
		odd := 0
		for j := i; j < len(nums); j++ {
			if _, ok := seen[nums[j]]; !ok {
				seen[nums[j]] = true;
				if nums[j] % 2 == 0 {
					even++;
				} else {
					odd++;
				}
			}
			if even == odd {
				longest = max(longest, j - i + 1)
			}
		}
	}
	return longest
}