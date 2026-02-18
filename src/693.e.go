package main

func hasAlternatingBits(n int) bool {
    last := n % 2;
    n /= 2;
    for n > 0 {
        if n % 2 == last {
            return false
        }
        last = n % 2
        n /= 2
    }
    return true
}