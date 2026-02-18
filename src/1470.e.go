package main

func shuffle(nums []int, n int) []int {
    var returnArray []int = make([]int, len(nums))
    for i := 0; i < n; i++ {
        returnArray[i*2] = nums[i]
    }
    j := 0
    for i := n; i < len(nums); i++ {
        returnArray[(j*2)+1] = nums[i];
        j++;
    }
    return returnArray
}