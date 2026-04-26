package main

func findLength(nums1 []int, nums2 []int) int {
	biggest := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			common := 0
			for i+common < len(nums1) && j+common < len(nums2) && nums1[i+common] == nums2[j+common] {
				common++;
			}
			biggest = max(common, biggest)
		}
	}
	return biggest
}