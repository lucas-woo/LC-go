package main

func bitwiseComplement(n int) int {
    if n == 0 {
        return 1
    }
    arr := make([]int, 0)
    for n > 0 {
        t := 0
        if n % 2 == 0 {
            t = 1
        }
        n /= 2
        arr = append(arr, t)
    }
    rN := 0;
    for i := len(arr) - 1; i >= 0; i-- {
        rN *= 2;
        rN += arr[i];
    }
    return rN
}