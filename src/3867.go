package main


import "sort"

func gcdSum(nums []int) int64 {
    prefixGcd := make([]int, len(nums))
    mx := nums[0]
    for i := 0; i < len(nums); i++ {
        mx = max(mx, nums[i])
        prefixGcd[i] = gcd(nums[i], mx)
    }

    sort.Ints(prefixGcd)
    i := 0;
    j := len(nums) - 1;
    var big int64
    for i < j {
        big += int64(gcd(prefixGcd[i], prefixGcd[j]))
        i++;
        j--;
    }
    return big
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
