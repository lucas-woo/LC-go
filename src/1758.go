package main

func minOperations(s string) int {
    ones := 0;
    zero := 0;
    for i := 0; i < len(s); i++ {
        if i & 1 == 1 {
            if s[i] == '1' {
                ones++;
            } else {
                zero++;
            }
        } else {
            if s[i] == '0' {
                ones++;
            } else {
                zero++;
            }
        }
    }
    return min(ones,zero)
}