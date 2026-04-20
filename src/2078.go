package main

func maxDistance(colors []int) int {
	left := 0;
	right := len(colors) - 1;

	biggest := 0
	for left <= right {
		if colors[left] != colors[right] {
			biggest = right - left
			break
		}
		right--;
	}
	left = 0;
	right = len(colors) - 1;
	for left <= right {
		if colors[left] != colors[right] {
			biggest = max(right - left, biggest)
			break
		}
		left++
	}	
	return biggest
}