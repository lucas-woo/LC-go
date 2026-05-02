package main

func maxRotateFunction(nums []int) int {
    n := len(nums)
    
    var sum int64 = 0
    var f int64 = 0
    
    for i := 0; i < n; i++ {
        sum += int64(nums[i])

        f += int64(i) * int64(nums[i])
    }
    
    total := f
    
    for k := 1; k < n; k++ {
        f = f + sum - int64(n)*int64(nums[n-k])
        
        if f > total {
            total = f
        }
    }
    
    return int(total)
}