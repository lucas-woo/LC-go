package main

func maxDistance(nums1 []int, nums2 []int) int {
	biggest := 0
	i := 0;
	j := 0;

	for i < len(nums1) && j < len(nums2) {
		if i > j || nums1[i] > nums2[j] {
			i++
			j = i
			continue
		}
		biggest = max(j - i, biggest)
		j++
	}
	return biggest 
}