package main

func maxArea(height []int) int {
	maxedArea := 0;
	left := 0;
	right := len(height) - 1;

	for i :=0 ; i <len(height); i++ {
		temp := ((right - left) * min(height[left], height[right]));
		if temp > maxedArea {
			maxedArea = temp
		}
		if height[left] > height[right] {
			right--;
		} else {
			left++;
		}
	}
	return maxedArea
}