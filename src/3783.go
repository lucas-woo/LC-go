package main

func mirrorDistance(n int) int {
    aCopy := n
    reversed := 0
    for n != 0 {
        reversed *= 10;
        reversed += n % 10;
        n /= 10;
    }
    return abs(aCopy - reversed)
}

func abs (n int) int {
    if n < 0 {
        return -n
    }
    return n
}