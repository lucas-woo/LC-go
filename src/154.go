package main

func findMin(nums []int) int {
    index := findMinIndex(nums)
    return nums[index]
}

func findMinIndex(arr []int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start < end {
        if arr[start] == arr[mid] && arr[end] == arr[mid] {
            start++
            end--
            continue
        }
		if mid > 0 && arr[mid] < arr[mid-1] {
			return mid
		} else if arr[end] >= arr[mid] {
			end = mid - 1 
		} else {
			start = mid + 1 
		}

		mid = start + (end-start)/2
	}

	return start
}
