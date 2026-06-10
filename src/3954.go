package main

func sumOfGoodIntegers(n int, k int) int {
    s := 0
    for i := 0; i <= n + k; i++ {
        if abs(n - i) <= k && n & i == 0 {
            s += i  
        }
    }
    return s
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}