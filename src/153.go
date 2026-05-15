package main
func findMin(nums []int) int {
    smallest := nums[0];
    for i := 1; i < len(nums); i++ {
        smallest = min(smallest, nums[i])
    }
    return smallest
}